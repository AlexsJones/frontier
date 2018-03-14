package processing

import (
	"fmt"

	"cloud.google.com/go/storage"
)

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
		quit:       make(chan bool),
	}
}

func objectURL(objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", objAttrs.Bucket, objAttrs.Name)
}

//Start ...
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:

				//Job has a hinged function on its struct which also is how it knows how to process
				//This lets us keep the worker generic...
				job.Process(job)

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
