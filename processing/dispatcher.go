package processing

import (
	"log"
	"sync"
)

var (
	maxWorkers = 10
)

//Dispatcher object contains Jobs
type Dispatcher struct {
	WorkerPool chan chan Job
	//JobQueue is pushed into to with new Job structs
	JobQueue chan Job
}

//Run will spawn the dispatch function loop in a new routine
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
		log.Printf("Started worker %d\n", i)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool
				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job) //<-----Enter the async function here
		}
	}
}

var instance *Dispatcher
var once sync.Once

//GetDispatcher is a singleton
func GetDispatcher() *Dispatcher {
	once.Do(func() {
		instance = &Dispatcher{WorkerPool: make(chan chan Job, maxWorkers), JobQueue: make(chan Job)}
	})
	return instance
}
