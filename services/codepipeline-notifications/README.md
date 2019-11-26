[![Serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)

# CodePipeline Notifications

A serverless service for reporting on Codepipeline and CodeBuild events to a Slack Channel.

Will report account wide so deploy in any accounts you want to report on.

## Dependencies

* DynamoDB (Inside AWS)
  - Used to store the Slack Message Identifier against the Build Identifier so it can continuously update a single message. 


## Deployment

```zsh
# Change region as appropriate.
sls deploy --region ap-southeast-2
```

### Required Environment Variables

```zsh
# Global:
export AWS_SLACK_NOTIFICATIONS_OAUTH_ACCESS_TOKEN=xoxp-825209819346-837534187124-837535618052-6597eb2eaceccd85340e0fe5033b43db

# Service Specific:
export CODEPIPELINE_DEPLOYMENT_NOTIFICATIONS_SLACK_CHANNEL=CQPR31LRM
```

### Required OAuth Scopes (Slack)

```zsh
chat:write:bot # Send messages as Application
```