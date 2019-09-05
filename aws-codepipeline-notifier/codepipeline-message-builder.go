package main

import (
	"fmt"
	"github.com/whithajess/aws-codepipeline-slack-notifications/internal/shared/message"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nlopes/slack"
)

// # https://docs.aws.amazon.com/codepipeline/latest/userguide/detect-state-changes-cloudwatch-events.html
// # https://docs.aws.amazon.com/codebuild/latest/APIReference/API_BuildPhase.html

func BuildAndSendSlackMessage(detail events.CodePipelineEventDetail, buildID string) {
	fmt.Println("## BuildSlackMessage:")
	var blocks []slack.Block // Full Set of blocks that make up the slack message.

	// TODO: Top Message

	var detailTextBlocks []*slack.TextBlockObject

	DetailTextBlockBuilder(&detailTextBlocks, detail)

	textInfo := slack.NewTextBlockObject("mrkdwn", "*Pipeline Status:*", false, false)

	blocks = append(blocks, slack.NewSectionBlock(textInfo, detailTextBlocks, nil, slack.SectionBlockOptionBlockID("first_phase_block")))

	channelID := os.Getenv("CodePipelineSlackChannelID")

	message.CreateOrUpdateMessage(channelID, buildID, blocks, slack.Attachment{})

	// TODO: Log links etc.
}

func DetailTextBlockBuilder(phaseTextBlocks *[]*slack.TextBlockObject, detail events.CodePipelineEventDetail ) {
	statusIconMapping := map[string]string{
		"": 									   message.BuildPhasesUnknown,
		string(events.CodePipelineStateStarted):   message.BuildPhasesInProgress,
		events.CodePipelineStateSucceeded:         message.BuildPhasesSucceeded,
		events.CodePipelineStateResumed:           message.BuildPhasesInProgress,
		events.CodePipelineStateFailed:            message.BuildPhasesFailed,
		events.CodePipelineStateCanceled:          message.BuildPhasesStopped,
		events.CodePipelineStateSuperseded:        message.BuildPhasesStopped,
	}

	phaseTextBlock := slack.NewTextBlockObject(
			"plain_text",
		    fmt.Sprintf("%s %s", string(statusIconMapping[string(detail.State)]), string(detail.Stage)),
			true,
			false,
	)
	*phaseTextBlocks = append(*phaseTextBlocks, phaseTextBlock)
}