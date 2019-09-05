module github.com/whithajess/aws-codepipeline-slack-notifications

go 1.12

require (
	github.com/aws/aws-lambda-go v1.13.0
	github.com/aws/aws-sdk-go v1.23.16
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/nlopes/slack v0.5.1-0.20190823181507-4411ac942311
	github.com/pkg/errors v0.8.1 // indirect
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
)

replace github.com/aws/aws-lambda-go => github.com/whithajess/aws-lambda-go v1.13.2-0.20191030023142-ba8d4131ff69
