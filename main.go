package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if request.RequestContext.HTTP.Method == http.MethodGet && request.RawPath == "/hello" {
		return events.APIGatewayV2HTTPResponse{
			Body:       "Hello, World!",
			StatusCode: http.StatusOK,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		Body:       "Not found",
		StatusCode: http.StatusNotFound,
	}, nil
}

func main() {
	lambda.Start(handler)
}
