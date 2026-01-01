
Example of a very simple Go webserver running in a cloud vendor

Originally from <https://blog.john-pfeiffer.com/go-faas-with-aws-lambda/>

# Local Dev

export GOOGLE_CLOUD_PROJECT=<GCP_PROJECT_ID>

## Build

    go mod init github.com/johnpfeiffer/aws-go-lambda
    go mod tidy
    go build

*compiles a binary file "aws-go-lambda" that executes on apple silicon, for linux: GOOS=linux GOARCH=amd64 go build -v*

### To run locally
`export GOOGLE_CLOUD_PROJECT="example-id"`
 `./aws-go-lambda`

# Testing

`go test -v`

    === RUN   TestHandler
    --- PASS: TestHandler (0.00s)
    === RUN   TestGenericHandler_HTTP
    --- PASS: TestGenericHandler_HTTP (0.00s)
    PASS
    ok  	github.com/johnpfeiffer/aws-go-lambda	0.218s

`curl localhost:8080`

`curl -s -X POST localhost:8080 -H 'Content-Type: application/json' -d '{"value":"world"}'`

# Deploy

If built into a .zip can be uploaded to AWS S3 and setup in an AWS Lambda

Otherwise just leverages GitHub and Google Cloud integration to auto deploy

