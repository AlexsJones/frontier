#fetch some data from an s3 bucket
curl -X POST -H "X-HEADER-REGION: eu-west-1" -H "X-HEADER-BUCKET: testing.beamery.com" http://localhost:8080/api/v1/bucketfetch -d '{}'
