module github.com/RIPGlobal/aws-slack-notifications/services/guardduty-notifications

go 1.13

require (
	github.com/RIPGlobal/aws-slack-notifications v0.0.0-20191115032350-6ea14371ef2d
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go v1.25.35
	github.com/nlopes/slack v0.6.0
	github.com/stretchr/testify v1.4.0
)

// replace also can be used to inform the go tooling of the relative or absolute on-disk location of modules in a multi-module project, such as:
// replace example.com/project/foo => ../foo
// we use this to include internal packages outside the root of this service.
replace github.com/RIPGlobal/aws-slack-notifications => ../../
