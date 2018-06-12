package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyRequest demonstrates an input value
type MyRequest struct {
	Value string `json:"value"`
}

// MyResponse helps illustrate how AWS Lambda auto
type MyResponse struct {
	Message string `json:"message"`
	Created string `json:"created"`
}

// HandleRequest https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-handler-types.html
func HandleRequest(ctx context.Context, req MyRequest) (MyResponse, error) {
	t := time.Now().UTC()
	return MyResponse{
		Message: fmt.Sprintf("hi %s", req.Value),
		Created: fmt.Sprintf("%s", t.Format(time.RFC3339))}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
