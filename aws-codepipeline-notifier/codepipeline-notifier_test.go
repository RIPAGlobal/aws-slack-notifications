package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCodePipelineNotifier(t *testing.T) {
	data, err := ioutil.ReadFile("../internal/test-data/codepipeline-action-execution-stage-change-event.json")
	require.NoError(t, err)

	codePipelineEvent := events.CodePipelineEvent{}
	json.Unmarshal(data, &codePipelineEvent)

	CodePipelineNotifier(codePipelineEvent)
}