package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCodeBuildNotifier(t *testing.T) {
	data, err := ioutil.ReadFile("../../../internal/test-data/codebuild-phase-change-notification.json")
	require.NoError(t, err)

	codeBuildEvent := events.CodeBuildEvent{}
	json.Unmarshal(data, &codeBuildEvent)

	CodeBuildNotifier(codeBuildEvent)
}
