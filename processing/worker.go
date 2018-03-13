package processing

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"github.com/AlexsJones/frontier/components/google"
	uuid "github.com/satori/go.uuid"
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
				//----------------------------------------------------------------------
				ctx := context.TODO()
				// or error handling
				u2, err := uuid.NewV4()
				if err != nil {
					fmt.Printf("Something went wrong: %s", err)
					return
				}
				obj := google.GetBucket().Object("test-" + u2.String())
				w := obj.NewWriter(ctx)
				rdr := bytes.NewBuffer(job.Payload)
				dec := base64.NewDecoder(base64.StdEncoding, rdr)
				if _, err := io.Copy(w, dec); err != nil {
					log.Fatal(err.Error())
				}
				if err := w.Close(); err != nil {
					log.Fatal(err.Error())
				}

				// if public {
				// 	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
				// 		return nil, nil, err
				// 	}

				objAttrs, _ := obj.Attrs(ctx)
				log.Printf("URL: %s", objectURL(objAttrs))
				log.Printf("Size: %d", objAttrs.Size)
				log.Printf("MD5: %x", objAttrs.MD5)
				log.Printf("objAttrs: %+v", objAttrs)
				//----------------------------------------------------------------------
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
