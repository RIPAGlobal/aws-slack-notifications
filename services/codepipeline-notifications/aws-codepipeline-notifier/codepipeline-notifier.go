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
func CodePipelineNotifier(event events.SQSEvent) {

	fmt.Println("EventRecord:")
	prettyJson, err := json.MarshalIndent(event.Records[0], "", "    ")
	if err != nil {
		fmt.Println("### JSON Indent Error:")
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s", prettyJson)

	var pipelineEvent events.CodePipelineEvent

	// Take the escaped Json String of the Event from the SQS Event so we can use it
	rawEscapedJsonString := []byte(event.Records[0].Body)
	rawJsonBody := (*json.RawMessage)(&rawEscapedJsonString)

	err = json.Unmarshal(*rawJsonBody, &pipelineEvent)
	if err != nil {
		fmt.Println("### JSON Unmarshal Error:")
		fmt.Println(err.Error())
	}

	fmt.Println("## CodePipelineEvent")
	fmt.Printf("\n## EventID: %s", pipelineEvent.ID)
	fmt.Printf("\n## Execution: %s", pipelineEvent.Detail.ExecutionId)

	// The default should never get hit with the settings in the serverless.yml for CloudWatch detail types to listen to
	// Useful for debugging and showing whats supported.
	//
	switch pipelineEvent.DetailType {
	case events.CodePipelineExecutionEventDetailType, events.CodePipelineActionEventDetailType, events.CodePipelineStageEventDetailType:
		executionID := pipelineEvent.Detail.ExecutionId
		BuildAndSendSlackMessage(pipelineEvent.Detail, executionID)
	default:
		fmt.Printf("\n# Non-Matched Event - Do nothing")
	}
}

func main() {
	lambda.Start(CodePipelineNotifier)
}
