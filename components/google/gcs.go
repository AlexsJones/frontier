package google

import (
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

var (
	instance   *storage.BucketHandle
	once       sync.Once
	projectID  = os.Getenv("GOOGLE_PROJECT")
	bucketName = os.Getenv("GOOGLE_BUCKET")
)

//GetBucket returns the handle for operating on the gcs bucket
func GetBucket() *storage.BucketHandle {
	once.Do(func() {

		if bucketName == "" || projectID == "" {
			fmt.Println("Requires GOOGLE_PROJECT and GOOGLE_BUCKET for the example to work.")
			os.Exit(1)
		}
		ctx := context.Background()
		// Creates a client.
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		// Creates a Bucket instance.
		instance = client.Bucket(bucketName)
		// Creates the new bucket.
		if err := instance.Create(ctx, projectID, nil); err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	})
	return instance
}
