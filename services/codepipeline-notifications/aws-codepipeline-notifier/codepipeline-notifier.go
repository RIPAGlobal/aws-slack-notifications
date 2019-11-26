package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SlackMessageIdentifier struct {
	BuildID string `json:"buildID"`
	TS      string `json:"ts"`
}

// CodePipelineNotifier is the lambda handler invoked by the `lambda.Start` function call
func CodePipelineNotifier(event events.CodePipelineEvent) {
	fmt.Println("### CodePipelineEvent")
	fmt.Printf("\n### EventID: %s", event.ID)
	fmt.Printf("\n### Execution: %s", event.Detail.ExecutionId)
	fmt.Println("EventFull:")
	prettyJson, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		fmt.Println("## JSON Indent Error:")
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s", prettyJson)

	// The default should never get hit with the settings in the serverless.yml for CloudWatch detail types to listen to
	// Useful for debugging and showing whats supported.
	//
	switch event.DetailType {
	case events.CodePipelineExecutionEventDetailType, events.CodePipelineActionEventDetailType, events.CodePipelineStageEventDetailType:
		executionID := event.Detail.ExecutionId
		BuildAndSendSlackMessage(event.Detail, executionID)
	default:
		fmt.Printf("\n# Non-Matched Event - Do nothing")
	}
}

// Generic Error printer if exists - not much else we can do with them.
func HandleErrors(err error, event events.CodePipelineEvent) {
	if err != nil {
		fmt.Println("# Error:")
		fmt.Println(err.Error())
		fmt.Println("# Request:")
		prettyJson, err := json.MarshalIndent(event, "", "    ")
		if err != nil {
			fmt.Println("# JSON Indent Error:")
			fmt.Println(err.Error())
		}
		fmt.Printf("\n%s", prettyJson)
	}
}

func main() {
	lambda.Start(CodePipelineNotifier)
}
