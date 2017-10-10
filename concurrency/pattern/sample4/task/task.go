package task

import "sync"

type Worker interface {
	Work()
}

type Task struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGourutine int) *Task {

	t := Task{
		work: make(chan Worker),
	}

	t.wg.Add(maxGourutine)

	for i := 0; i < maxGourutine; i++ {
		go func() {

			for w := range t.work {
				w.Work()
			}
			t.wg.Done()
		}()
	}

	return &t
}

func (t *Task) ShutDown() {
	close(t.work)
	t.wg.Wait()
}

func (t *Task) Do(w Worker) {
	t.work <- w
}
