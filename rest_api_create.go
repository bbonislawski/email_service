package main

import (
  "fmt"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  repo := EmailsRepository{}
  err := repo.Create(request.Body)

  if err != nil {
    return events.APIGatewayProxyResponse{
      Body: fmt.Sprintf(`{"message": "%s"}`, err.Error()),
      StatusCode: 422,
    }, nil
  } else {
    return events.APIGatewayProxyResponse{
      Body: `{"message": "Email sent!"}`,
      StatusCode: 200,
    }, nil
  }
}

func main() {
  lambda.Start(Handler)
}
