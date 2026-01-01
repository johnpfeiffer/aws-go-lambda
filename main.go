package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyRequest demonstrates an input value from the inbound request
type MyRequest struct {
	Value string `json:"value"`
}

// MyResponse structures the output response as JSON
type MyResponse struct {
	Message string `json:"message"`
	Created string `json:"created"`
}

// generateResponse is an example function
func generateResponse(inputValue string) (MyResponse, error) {
	t := time.Now().UTC()
	return MyResponse{
		Message: fmt.Sprintf("hi %s", inputValue),
		Created: fmt.Sprintf("%s", t.Format(time.RFC3339))}, nil
}

// HandleRequest for an AWS Lambda https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
func HandleRequest(ctx context.Context, req MyRequest) (MyResponse, error) {
	return generateResponse(req.Value)
}

// StartServer with routes from https://pkg.go.dev/net/http#ServeMux
func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/{path...}", GenericHandler)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// GenericHandler for HTTP requests
func GenericHandler(w http.ResponseWriter, r *http.Request) {
	var req MyRequest
	if r.Body != nil {
		defer r.Body.Close()
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&req); err != nil && err != io.EOF {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := generateResponse(req.Value)
	if err != nil {
		http.Error(w, "Error encountered", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error encountered", http.StatusInternalServerError)
	}
}

func main() {

	// https://github.com/GoogleCloudPlatform/cloud-run-microservice-template-go
	if os.Getenv("GOOGLE_CLOUD_PROJECT") != "" {
		port := os.Getenv("PORT")
		if port == "" {
			port = ":8080"
		}
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
		log.Printf("listening on port %s", port)
		projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
		log.Printf("running in Google Cloud Project %s", projectID)
		StartServer(port)
	}

	// AWS Lambda
	lambda.Start(HandleRequest)
}
