package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Hello World",
	}

	return response, nil
}
