package main

import (
  "context"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/rubyist/circuitbreaker"
)

type Adapter interface {
  SendEmail(to, subject, content string) (string, error)
}

func Handler(ctx context.Context, e events.DynamoDBEvent) {
  for _, record := range e.Records {
    if record.EventName == "INSERT" {
      to := record.Change.NewImage["Address"].String()
      subject := record.Change.NewImage["Subject"].String()
      content := record.Change.NewImage["Content"].String()
      sendEmail(to, subject, content)
    }
  }
}

func sendEmail(to, subject, content string) (string, error) {
  var a Adapter = MailgunAdapter{}

  cb := circuit.NewThresholdBreaker(3)

  for {
    if cb.Ready() {
      result, err := a.SendEmail(to, subject, content)

      if err != nil {
        cb.Fail()
        continue
      }
      cb.Success()
      return result, nil
    } else {
      a = SendgridAdapter{}

      result, err := a.SendEmail(to, subject, content)
      if err != nil {
        return result, err
      } else {
        return result, nil
      }
    }
  }

}

func main() {
  lambda.Start(Handler)
}
