[![Serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)

# GuardDuty Slack Reporter

A serverless service for reporting findings from GuardDuty via CloudWatch events to a Slack Channel.

***Suggested that this is installed into a Security/Audit Account - where you house your GuardDuty (master)***

See Resources for why and how:
* [Your single AWS account is a serious risk](https://cloudonaut.io/your-single-aws-account-is-a-serious-risk/)
* [You need more than one AWS account: AWS bastions and assume-role](https://blog.coinbase.com/you-need-more-than-one-aws-account-aws-bastions-and-assume-role-23946c6dfde3)
* [AWS security configurations and best practices - GuardDuty](https://asecure.cloud/l/s_guardduty/)

Otherwise if you have a single account or multiple accounts without a master deploy individually in each.

## Deployment



### Environment Variables

```zsh
# Global:
export AWS_SLACK_NOTIFICATIONS_OAUTH_ACCESS_TOKEN=xoxp-825209819346-837534187124-837535618052-6597eb2eaceccd85340e0fe5033b43db

# Service Specific:
export GUARD_DUTY_NOTIFICATIONS_SLACK_CHANNEL=CQN0636KX
```

### Required OAuth Scopes (Slack)

```zsh
chat:write:bot # Send messages as Application
```