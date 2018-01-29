#!/bin/bash

aws ssm put-parameter --name SENDGRID_API_KEY --value YOUR_SENDGRID_API_KEY --type String --overwrite
aws ssm put-parameter --name MG_API_KEY --value YOUR_MAILGUN_API_KEY --type String --overwrite
aws ssm put-parameter --name MG_PUBLIC_API_KEY --value YOUR_MAILGUN_PUBLIC_API_KEY --type String --overwrite
aws ssm put-parameter --name MG_URL --value https://api.mailgun.net/v3 --type String --overwrite
aws ssm put-parameter --name MG_DOMAIN --value YOUR_DOMAIN_FOR_SENDING_EMAILS --type String --overwrite
aws ssm put-parameter --name API_KEY_TOKEN --value YOUR_TOKEN_TO_CONTROL_ACCESS --type String --overwrite
