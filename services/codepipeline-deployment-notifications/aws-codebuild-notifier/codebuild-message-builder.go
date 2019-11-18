package main

import (
	"fmt"
	"github.com/RIPGlobal/aws-slack-notifications/internal/shared/message"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nlopes/slack"
)

// # https://docs.aws.amazon.com/codepipeline/latest/userguide/detect-state-changes-cloudwatch-events.html

func BuildAndSendSlackMessage(event events.CodeBuildEvent, buildID string) {
	fmt.Println("## BuildSlackMessage:")
	// If this event has come from Codepipeline assume we do not need to add the Source Details
	if event.Detail.AdditionalInformation.Source.Type != "CODEPIPELINE" {
		log.Fatalln("Currently only taking CodePipeline messages")
	}

	var blocks []slack.Block // Full Set of blocks that make up the slack message.

	var phaseTextBlocks []*slack.TextBlockObject
	PhaseTextBlockBuilder(&phaseTextBlocks, event.Detail.AdditionalInformation.Phases)

	textInfo := slack.NewTextBlockObject("mrkdwn", "*Build Status:*", false, false)

	// Deal with the case of over 10 message.
	//  - This will break over 20 but currently impossible
	//
	if len(phaseTextBlocks) > 10 {
		blocks = append(blocks, slack.NewSectionBlock(textInfo, phaseTextBlocks[:10], nil, slack.SectionBlockOptionBlockID("first_phase_block")))
		blocks = append(blocks, slack.NewSectionBlock(nil, phaseTextBlocks[10:], nil, slack.SectionBlockOptionBlockID("second_phase_block")))
	} else {
		blocks = append(blocks, slack.NewSectionBlock(textInfo, phaseTextBlocks, nil, slack.SectionBlockOptionBlockID("first_phase_block")))
	}

	// This is the old way of using Buttons - but seems to still be the official advice to use this for "Button Links"
	// otherwise you essentially have to build a listener to process the requests / just use text links
	// TODO: Revisit this in the future.
	action := slack.AttachmentAction{
		Name:  "Build Logs",
		Text:  "Build Logs",
		Style: "", // TODO: Should change based on Status
		Type:  "button",
		URL:   event.Detail.AdditionalInformation.Logs.DeepLink,
	}

	// Package up the button as an Attachment
	attachment := slack.Attachment{
		Actions: []slack.AttachmentAction{action},
	}

	channelID := os.Getenv("CODEPIPELINE_DEPLOYMENT_NOTIFICATIONS_SLACK_CHANNEL")

	message.CreateOrUpdateMessage(channelID, buildID, blocks, attachment)
}

func PhaseTextBlockBuilder(phaseTextBlocks *[]*slack.TextBlockObject, phases []events.CodeBuildPhase) {
	statusIconMapping := map[string]string{
		"": message.BuildPhasesUnknown,
		string(events.CodeBuildPhaseStatusFailed): message.BuildPhasesFailed,
		events.CodeBuildPhaseStatusFault:          message.BuildPhasesFault,
		events.CodeBuildPhaseStatusQueued:         message.BuildPhasesQueued,
		events.CodeBuildPhaseStatusInProgress:     message.BuildPhasesInProgress,
		events.CodeBuildPhaseStatusStopped:        message.BuildPhasesStopped,
		events.CodeBuildPhaseStatusSucceeded:      message.BuildPhasesSucceeded,
		events.CodeBuildPhaseStatusTimedOut:       message.BuildPhasesTimedOut,
	}

	for _, phase := range phases {
		// "COMPLETED" doesn't seem to send a phase status assume its a success when received.
		if phase.PhaseType == events.CodeBuildPhaseTypeCompleted {
			phaseTextBlock := slack.NewTextBlockObject(
				"plain_text",
				fmt.Sprintf("%s %s", message.BuildPhasesSucceeded, string(phase.PhaseType)),
				true,
				false,
			)
			*phaseTextBlocks = append(*phaseTextBlocks, phaseTextBlock)
		} else {
			phaseTextBlock := slack.NewTextBlockObject(
				"plain_text",
				fmt.Sprintf("%s %s", string(statusIconMapping[string(phase.PhaseStatus)]), string(phase.PhaseType)),
				true,
				false,
			)
			*phaseTextBlocks = append(*phaseTextBlocks, phaseTextBlock)
		}
	}
}
