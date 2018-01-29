package main

import (
  "time"
  "os"
  "errors"
  "encoding/json"
  "github.com/satori/go.uuid"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
type EmailsRepository struct { }

type EmailData struct {
  EmailId string`json:"EmailId"`
  Content string`json:"Content"`
  Subject string`json:"Subject"`
  Address string`json:"Address"`
  EmailTime string`json:"EmailTime"`
}


var (
  ErrEmailNotProvided = errors.New("no Address was provided in the HTTP body")
  ErrSubjectNotProvided = errors.New("no Subject was provided in the HTTP body")
  ErrContentNotProvided = errors.New("no Content was provided in the HTTP body")
)

func service() (*dynamodb.DynamoDB) {
  sess, _ := session.NewSession(&aws.Config{
    Region: aws.String(os.Getenv("REGION"))},
  )
  return dynamodb.New(sess)
}

func(e EmailsRepository) All() ([]EmailData) {
  var emails []EmailData
  svc := service()

  params := &dynamodb.ScanInput{
    TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
  }

  result, _ := svc.Scan(params)

  for _, i := range result.Items {
    item := EmailData{}
    _ = dynamodbattribute.UnmarshalMap(i, &item)
    emails = append(emails, item)
  }

  return emails
}

func(e EmailsRepository) Create(requestBody string) (error) {
  svc := service()

  emailData := EmailData{}
  json.Unmarshal([]byte(requestBody), &emailData)
  emailData.EmailId = uuid.Must(uuid.NewV4()).String()
  emailData.EmailTime = time.Now().Format(time.RFC3339)

  switch {
  case len(emailData.Subject) < 1:
    return ErrSubjectNotProvided
  case len(emailData.Content) < 1:
    return ErrSubjectNotProvided
  case len(emailData.Address) < 1:
    return ErrEmailNotProvided
  }

  av, err := dynamodbattribute.MarshalMap(emailData)

  input := &dynamodb.PutItemInput{
    Item: av,
    TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
  }

  _, err = svc.PutItem(input)
  return err
}
