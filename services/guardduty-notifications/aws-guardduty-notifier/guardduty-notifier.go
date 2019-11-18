package main

import (
	"encoding/json"
	"github.com/RIPGlobal/aws-slack-notifications/internal/shared/message"
	"os"

	"github.com/nlopes/slack"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// GuardDutyFinding - Takes relevant info out of the json payload from GuardDuty
type GuardDutyFinding struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Severity    json.Number `json:"severity"`
	Type        string      `json:"type"`
	AccountId   string      `json:"accountId"`
}

// Reporter - Listens for CloudWatch events of GuardDuty Findings
// Then formats these and sends them to Slack
func GuardDutyNotifier(event events.CloudWatchEvent) {
	// TODO: Worth moving this into aws-lambda-go events
	var finding GuardDutyFinding
	json.Unmarshal([]byte(event.Detail), &finding)

	blocks := []slack.Block{
		slack.NewSectionBlock(
			slack.NewTextBlockObject("mrkdwn", "*GuardDuty Finding:*", false, false),
			nil,
			nil,
			slack.SectionBlockOptionBlockID("second_phase_block")),
	}

	attachment := slack.Attachment{
		Title: finding.Title,
		Text:  finding.Description,
		Color: "danger",
		Fields: []slack.AttachmentField{
			{
				Title: "Account ID",
				Value: finding.AccountId,
			},
			{
				Title: "Severity",
				Value: string(finding.Severity),
			},
			{
				Title: "Type",
				Value: finding.Type,
			},
		},
	}

	channelID := os.Getenv("GUARD_DUTY_NOTIFICATIONS_SLACK_CHANNEL")

	message.CreateMessage(channelID, blocks, attachment)
}

func main() {
	lambda.Start(GuardDutyNotifier)
}
