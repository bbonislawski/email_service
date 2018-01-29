#!/bin/bash

GOOS=linux go build -o bin/email_processing main.go mock_adapter.go mailgun_adapter.go sendgrid_adapter.go

GOOS=linux go build -o bin/rest_api_index rest_api_index.go emails_repository.go
GOOS=linux go build -o bin/rest_api_create rest_api_create.go emails_repository.go
