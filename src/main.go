package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//Handler - Is called after main
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	t := fmt.Sprintf("Time is %v", time.Now()) //fmt.Sprint returns a string, which is assigned to 't'

	return events.APIGatewayProxyResponse{Body: t, StatusCode: 200}, nil //Respond with 't' and a 200 Code
}

func main() {
	lambda.Start(Handler)
}
