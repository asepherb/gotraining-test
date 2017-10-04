package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type client struct {
	name   string
	reader *bufio.Reader
}

func (c *client) TypeAsContext() {

	for {

		line, err := c.reader.ReadString('\n')
		if err != nil {

			switch e := err.(type) {
			case *net.OpError:
				if !e.Temporary() {
					log.Println("temporary: client leave chat")
					return
				}

			case *net.AddrError:
				if !e.Temporary() {
					log.Println("temporary: client leave chat")
					return
				}

			case *net.DNSConfigError:
				if !e.Temporary() {
					log.Println("temporary: client leave chat")
					return
				}

			default:
				if err == io.EOF {
					log.Println("EOF, client leave chat")
					return
				}

				log.Println("read-routine", err)
			}
		}

		fmt.Println(line)
	}
}

type temporary interface {
	Temporary() bool
}

func (c *client) BehaviorAsContext() {
	for {

		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case temporary:
				if !e.Temporary() {
					log.Println("temporary : client leaving chat")
					return
				}

			default:
				if err == io.EOF {
					log.Println("EOF, client leave chat")
					return
				}

				log.Println("read routine", err)
			}
		}

		fmt.Println(line)
	}
}

func main() {

}
