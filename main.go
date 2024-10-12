package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == http.MethodGet && request.Path == "/hello" {
		return events.APIGatewayProxyResponse{
			Body:       "Hello, World!",
			StatusCode: http.StatusOK,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Not found",
		StatusCode: http.StatusNotFound,
	}, nil
}

func main() {
	lambda.Start(handler)
}
