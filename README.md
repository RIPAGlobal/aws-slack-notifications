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

### CodePipeline Deployment Notifications

### GuardDuty Notifications

## Service Setup

### Slack

***1.*** Login to https://api.slack.com/apps and Create New App
   * The App name you use will be the user reporting the guard duty findings
   * The App will need to be installed for you to get an OAuth token.


***2.*** After creating the application in Slack we need to set tokens for the App to use:
   * These are region specific if you haven't set `AWS_PROFILE` you may want to use the flag `--region`
```bash
  # The verification token
  # can be found under Basic Information in the App on https://api.slack.com/apps
  # This gives us the ability to check the messages sent to the App are actually coming from Slack
  aws ssm put-parameter --name guardBotOAuthAccessToken --type String --value SecretToken

  # The channel we want to post into
  aws ssm put-parameter --name guardChannel --type String --value ChannelID
```

***3.*** After this is all set you will need to deploy the application
```bash
  make
  sls deploy
```

#### Required Permissions


### Environment Variables

These are service specific but global defaults:
```zsh
export AWS_SLACK_NOTIFICATIONS_OAUTH_ACCESS_TOKEN=
```
