package main

import (
  "log"
  "os"
  "gopkg.in/mailgun/mailgun-go.v1"
)
type MailgunAdapter struct { }

func(a MailgunAdapter) SendEmail(to, subject, content string) (string, error) {
  mg := mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"), os.Getenv("MG_PUBLIC_API_KEY"))
  message := mg.NewMessage(os.Getenv("EMAIL_FROM"), subject, content, to)
  message.SetHtml(content)
  resp, _, err := mg.Send(message)
  if err != nil {
    log.Fatal(err)
    return "", err
  } else {
    return resp, nil
  }
}

