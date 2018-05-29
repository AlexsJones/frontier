package aws

import (
	"fmt"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fatih/color"
)

var (
	instance *storage.BucketHandle
	once     sync.Once
)

//NewS3S SessionClient ...
func NewS3S(region string) (*s3.S3, error) {
	sess, err := session.NewSession()
	svc := s3.New(sess, aws.NewConfig().WithRegion(region))
	if err != nil {
		return nil, err
	}
	return svc, nil
}

//S3BucketList from AWS_BUCKET_LIST
func S3BucketKeyList(svc *s3.S3, bucket string) error {

	params := &s3.ListObjectsInput{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(100),
	}
	color.Yellow("Fetching object pages")
	err := svc.ListObjectsPages(params,
		func(page *s3.ListObjectsOutput, last bool) bool {
			for _, object := range page.Contents {
				fmt.Printf("%s:%s", *params.Bucket, *object.Key)
			}
			return true
		},
	)
	if err != nil {
		fmt.Println("Error listing", *params.Bucket, "objects:", err)
	}
	return nil
}
