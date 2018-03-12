package processing

//Dispatcher object contains Jobs
type Dispatcher struct {
	WorkerPool chan chan Job
}

//NewDispatcher for managing Jobs is created
func NewDispatcher(maxWorkers int) *Dispatcher {

	return &Dispatcher{WorkerPool: make(chan chan Job, maxWorkers)}
}
