package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// CodeBuildNotifier is the lambda handler invoked by the `lambda.Start` function call
//
func CodeBuildNotifier(event events.SQSEvent) {

	fmt.Println("EventRecord:")
	prettyJson, err := json.MarshalIndent(event.Records[0], "", "    ")
	if err != nil {
		fmt.Println("### JSON Indent Error:")
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s", prettyJson)

	var codebuildEvent events.CodeBuildEvent

	// Take the escaped Json String of the Event from the SQS Event so we can use it
	rawEscapedJsonString := []byte(event.Records[0].Body)
	rawJsonBody := (*json.RawMessage)(&rawEscapedJsonString)

	err = json.Unmarshal(*rawJsonBody, &codebuildEvent)
	if err != nil {
		fmt.Println("### JSON Unmarshal Error:")
		fmt.Println(err.Error())
	}

	fmt.Println("## CodeBuildEvent")
	fmt.Printf("\n## EventID: %s", codebuildEvent.ID)
	fmt.Printf("\n## BuildID: %s", codebuildEvent.Detail.BuildID)

	// The default should never get hit with the settings in the serverless.yml for CloudWatch detail types to listen to
	// Useful for debugging and showing whats supported.
	//
	switch codebuildEvent.DetailType {
	case events.CodeBuildStateChangeDetailType, events.CodeBuildPhaseChangeDetailType:
		buildID := codebuildEvent.Detail.BuildID
		BuildAndSendSlackMessage(codebuildEvent, buildID)
	default:
		fmt.Printf("\n# Non-Matched Event - Do nothing")
	}
}

func main() {
	lambda.Start(CodeBuildNotifier)
}
