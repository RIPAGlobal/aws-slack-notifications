package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCodePipelineNotifier(t *testing.T) {
	data, err := ioutil.ReadFile("../../../internal/test-data/codepipeline-action-execution-stage-change-event.json")
	require.NoError(t, err)

	message := events.SQSMessage{Body: string([]byte(data))}

	sqsEvent := events.SQSEvent{[]events.SQSMessage{message}}

	CodePipelineNotifier(sqsEvent)
}
