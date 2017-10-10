package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type temporary interface {
	Temporary() bool
}

type message struct {
	data string
	conn net.Conn
}

type client struct {
	name   string
	room   *Room
	reader *bufio.Reader
	writer *bufio.Writer
	wg     sync.WaitGroup
	conn   net.Conn
}

func (c *client) read() {
	for {

		//wait for message to arrive
		line, err := c.reader.ReadString('\n')

		if err == nil {
			c.room.outgoing <- message{
				data: line,
				conn: c.conn,
			}
			continue
		}

		if e, is := err.(temporary); is && !e.Temporary() {
			log.Println("temporary: client is leaving chat")
			c.wg.Done()
			return
		}

		if err == io.EOF {
			log.Println("EOF: client is leaving chat")
			c.wg.Done()
			return
		}

		log.Println("read-routine", err)
	}
}

func (c *client) write(m message) {
	msg := fmt.Sprintf("%s %s", c.name, m.data)
	log.Printf(msg)

	c.writer.WriteString(msg)
	c.writer.Flush()
}

func (c *client) drop() {
	c.conn.Close()
	c.wg.Wait()
}

func newClient(room *Room, conn net.Conn, name string) *client {

	c := client{
		name:   name,
		room:   room,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		conn:   conn,
	}

	c.wg.Add(1)
	go c.read()

	return &c
}

type Room struct {
	listener net.Listener
	clients  []*client
	joining  chan net.Conn
	outgoing chan message
	shutdown chan struct{}
	wg       sync.WaitGroup
}

func (r *Room) sendGroupMessage(m message) {
	for _, c := range r.clients {
		if c.conn != m.conn {
			c.write(m)
		}
	}
}

func (r *Room) join(conn net.Conn) {

	name := fmt.Sprintf("conn : %d", len(r.clients))
	log.Println("new client joining chat:", name)

	c := newClient(r, conn, name)
	r.clients = append(r.clients, c)
}

func (r *Room) start() {

	r.wg.Add(2)

	//chatroom processing
	go func() {
		for {
			select {
			case message := <-r.outgoing:
				r.sendGroupMessage(message)
			case conn := <-r.joining:
				r.join(conn)
			case <-r.shutdown:
				r.wg.Done()
				return
			}
		}
	}()

	//chatroom connection accept
	go func() {

		var err error

		if r.listener, err = net.Listen("tcp", ":6000"); err != nil {
			log.Fatalln(err)
		}

		log.Println("chat room started: 6000")
		for {
			conn, err := r.listener.Accept()
			if err != nil {

				if e, is := err.(temporary); is {
					if !e.Temporary() {
						log.Println("temporary: chat room is shutting down")
						r.wg.Done()
						return
					}
				}

				log.Println("accept-goroutine", err)
				continue
			}

			r.joining <- conn
		}
	}()
}

func (r *Room) Close() error {

	r.listener.Close()

	close(r.shutdown)

	r.wg.Wait()

	for _, c := range r.clients {
		c.drop()
	}
	return nil
}

func New() *Room {

	chatRoom := Room{
		joining:  make(chan net.Conn),
		outgoing: make(chan message),
		shutdown: make(chan struct{}),
	}

	chatRoom.start()

	return &chatRoom
}
