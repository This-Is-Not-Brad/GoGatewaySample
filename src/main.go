package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//Handler - Is called after main
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	now := time.Now()

	t := fmt.Sprintf("Time is %v", now)

	return events.APIGatewayProxyResponse{Body: t, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
