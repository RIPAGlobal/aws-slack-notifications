package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCodeBuildNotifier(t *testing.T) {
	data, err := ioutil.ReadFile("../../../internal/test-data/codebuild-phase-change-notification.json")
	require.NoError(t, err)

	message := events.SQSMessage{Body: string([]byte(data))}

	sqsEvent := events.SQSEvent{[]events.SQSMessage{message}}

	CodeBuildNotifier(sqsEvent)
}
