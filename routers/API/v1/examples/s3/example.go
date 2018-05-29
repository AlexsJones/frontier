package s3

import (
	"fmt"
	"net/http"

	"github.com/AlexsJones/frontier/components/aws"
	"github.com/AlexsJones/frontier/processing"
	"github.com/fatih/color"
)

//Post is for illustrative purposes of how to use the processing API
func Post(arg1 http.ResponseWriter, arg2 *http.Request) {

	region := arg2.Header.Get("X-HEADER-REGION")
	bucket := arg2.Header.Get("X-HEADER-BUCKET")

	if region == "" || bucket == "" {
		arg1.WriteHeader(http.StatusBadRequest)
		color.Red("Bad request")
		return
	}
	var j processing.Job

	j.Process = func(j processing.Job) {
		color.Yellow(fmt.Sprintf("Processing bucket %s %s \n", region, bucket))
		s3client, err := aws.NewS3S(region)
		if err != nil {
			color.Red("Error creating client")
			arg1.WriteHeader(http.StatusInternalServerError)
		}

		err = aws.S3BucketKeyList(s3client, bucket)
		if err != nil {
			color.Red("Error Listing bucket Keys")
			arg1.WriteHeader(http.StatusInternalServerError)
		}
	}
	processing.GetDispatcher().JobQueue <- j

	arg1.WriteHeader(http.StatusOK)
}
