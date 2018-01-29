package main

import (
  "log"
  "os"
  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
)
type SendgridAdapter struct { }

func(a SendgridAdapter) SendEmail(to, subject, content string) (string, error) {

  from := mail.NewEmail("", os.Getenv("EMAIL_FROM"))
  toMail := mail.NewEmail("", to)
  message := mail.NewSingleEmail(from, subject, toMail, content, content)
  client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
  response, err := client.Send(message)

  if err != nil {
    log.Println(err)

    return "", err
  } else {
    s := string(response.Body[:])
    return s, nil
  }
}

