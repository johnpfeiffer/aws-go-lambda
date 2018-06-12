package main

import (
	"context"
	"testing"
)

func TestHandler(t *testing.T) {
	var ctx context.Context
	response, err := HandleRequest(ctx, MyRequest{Value: "john"})
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	expected := "hi john"
	if response.Message != expected {
		t.Errorf("Incorrect response, expected: '%s' , but received: %s", expected, response.Message)
	}
}
