[![Serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/RIPGlobal/aws-slack-notifications)](https://goreportcard.com/report/github.com/RIPGlobal/aws-slack-notifications)
[![License](https://img.shields.io/github/license/ripglobal/aws-slack-notifications.svg)](LICENSE)

# AWS Slack Notifications

A generic set up for sending AWS notifications to Slack.

Uses:
* Golang
* Serverless

## Services

Notifications are set up as individual Serverless services, making them deployable individually. While sharing common code.

### [CodePipeline Notifications](services/codepipeline-notifications/README.md)

### [GuardDuty Notifications](services/guardduty-notifications/README.md)

## Service Setup

### AWS

[Serverless Guide to AWS Credentials Setup](https://serverless.com/framework/docs/providers/aws/guide/credentials/)

### Slack

***1.*** Login to https://api.slack.com/apps and Create New App
   * The App name you use will be the user reporting the notifications in slack
   * The App will need to be installed for you to get an OAuth token.

#### Required Permissions

These are service specific (see Service README) but commonly:
```zsh
chat:write:bot # Send messages as AWS Notifications
```

### Environment Variables

These are service specific (see Service README) but commonly:
```zsh
export AWS_SLACK_NOTIFICATIONS_OAUTH_ACCESS_TOKEN=xoxp-825209819346-837534187124-837535618052-6597eb2eaceccd85340e0fe5033b43db
```