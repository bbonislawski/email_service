package main

import (
  "context"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

type Adapter interface {
  SendEmail(to, subject, content string) (string, error)
}

func Handler(ctx context.Context, e events.DynamoDBEvent) {
}

func sendEmail(to, subject, content string) (string, error) {
  var a Adapter = MailgunAdapter{}
  result, err := a.SendEmail(to, subject, content)

  return result, err
}

func main() {
  lambda.Start(Handler)
}
