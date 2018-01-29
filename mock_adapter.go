package main

import "errors"
type MockAdapter struct { }
type FailedMockAdapter struct { }

func(a MockAdapter) SendEmail(to, subject, content string) (string, error) {
  return "Successfully sent!", nil
}

func(a FailedMockAdapter) SendEmail(to, subject, content string) (string, error) {
  return "", errors.New("Sending email failed.")
}

