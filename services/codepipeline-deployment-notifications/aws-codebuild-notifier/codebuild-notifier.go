package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// CodeBuildNotifier is the lambda handler invoked by the `lambda.Start` function call
//
func CodeBuildNotifier(event events.CodeBuildEvent) {
	fmt.Println("### CodeBuildEvent")
	fmt.Printf("\n### EventID: %s", event.ID)
	fmt.Printf("\n### BuildID: %s", event.Detail.BuildID)
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
	case events.CodeBuildStateChangeDetailType, events.CodeBuildPhaseChangeDetailType:
		buildID := event.Detail.BuildID
		BuildAndSendSlackMessage(event, buildID)
	default:
		fmt.Printf("\n# Non-Matched Event - Do nothing")
	}
}

// Generic Error printer if exists - not much else we can do with them.
//
func HandleErrors(err error, event events.CodeBuildEvent) {
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
	lambda.Start(CodeBuildNotifier)
}
