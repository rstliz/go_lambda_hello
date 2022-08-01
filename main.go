package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"Name"`
}

type Response struct {
	Message string `json:"Message:"`
}

func hello(request Request) (Response, error) {
	return Response{Message: fmt.Sprintf("Hello %s.", request.Name)}, nil
}

func main() {
	lambda.Start(hello)
}
