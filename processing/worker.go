package processing

import "fmt"

//Worker ...
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

//NewWorker ...
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

//Start ...
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case _ = <-w.JobChannel:
				// we have received a work request.
				fmt.Println("Received work request")
			case <-w.quit:

				return
			}
		}
	}()
}

//Stop running the worker
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
