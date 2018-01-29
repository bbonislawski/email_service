package main

import (
  "encoding/json"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  repo := EmailsRepository{}
  emails := repo.All()

  jsonResult, _ := json.Marshal(emails)

  return events.APIGatewayProxyResponse{
    Body:       string(jsonResult),
    StatusCode: 200,
  }, nil

}

func main() {
  lambda.Start(Handler)
}
