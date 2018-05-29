#fetch some data from an s3 bucket
curl -X POST -H "X-HEADER-REGION: SOMEREGION" -H "X-HEADER-BUCKET: COMPANY.com" http://localhost:8080/api/v1/bucketfetch -d '{}'
