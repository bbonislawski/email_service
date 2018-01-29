package main

import (
  "time"
  "os"
  "errors"
  "encoding/json"
)
type MockEmailsRepository struct { }

func(e MockEmailsRepository) All() ([]EmailData) {
  var emails []EmailData
  item := EmailData{}

  item.Address = "test@example.com"
  item.Content = "Content Test"
  item.Subject = "Test Subject"
  item.EmailId = uuid.Must(uuid.NewV4()).String()
  item.EmailTime = time.Now().Format(time.RFC3339)

  emails = append(emails, item)

  return emails
}

func(e MockEmailsRepository) Create(requestBody string) (error) {
  return nil
}
