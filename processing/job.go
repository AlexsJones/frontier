package processing

//Job is a DTO for a unit to be processed
type Job struct {
	DTO interface{}
	//This is a function for processing the Job
	Process func(Job)
}
