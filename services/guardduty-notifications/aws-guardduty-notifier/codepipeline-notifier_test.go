package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestGuardDutyNotifier(t *testing.T) {
	data, err := ioutil.ReadFile("../../../internal/test-data/guardduty-finding-event.json")
	require.NoError(t, err)

	cloudWatchEvent := events.CloudWatchEvent{}
	json.Unmarshal(data, &cloudWatchEvent)

	GuardDutyNotifier(cloudWatchEvent)
}
