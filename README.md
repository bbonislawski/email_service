# Serverless Golang service for sending emails with Mailgun and Sendgrid as alternative

This service should support 2 email sending services, so in case of failure of one service, the second one would be used
as an alternative.

To solve this problem I've created simple json API with 2 actions which are documented [HERE](https://app.swaggerhub.com/apis/bbdev/email-service/2018-01-29T14-08-08#/default/post_mails)

I've implemented SendEmail function with very basic implementation of Circuit Breaker pattern that tries to send email
with Mailgun and in case of fail(it tries 3 times) changes adapter to SendgridAdapter and sends email with Sendgrid. In
future, it could save information that one of services is down and force service to user other one.


I don't have any commercial experience with Golang nor Serverless/AWS Lambda that's why there are missing:
- Filtering emails on index api
- Automated tests(It's pretty hard case to test and I didn't have enough time to learn golang enough to do it)
- Frontend service(but there is swagger)

This "service" is scalable to the limits of AWS, so I guess it's one of most scalable ways to do it. Function for
processing sending of emails is called by DynamoDB trigger.

I can see a lot of room for improvement since when both services are down, then it would fail to send email(and didn't
notify user about it) but it could be pretty easily fixed by adding field `Completed` and implementing service that
would fetch "failed calls" and try to send them again.

I've prepared architecture the way that it's really easy to implement other services for sending emails by simply adding
new adapter.

This was built with Serverless framework on AWS Lambda and Golang.
## Quick Start

1. Download dependencies

```
sh setup_dependencies.sh
```

2. Edit and upload your credentials

Open `setup_credentials.sh` and provide your data
```
sh setup_credentials.sh
```

3. Build

```
sh build.sh
```

4. Deploy!

```
serverless deploy
```
