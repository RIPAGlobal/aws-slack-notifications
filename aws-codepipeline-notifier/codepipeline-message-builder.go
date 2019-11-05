package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/whithajess/aws-codepipeline-slack-notifications/internal/shared/message"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/nlopes/slack"
)

// # https://docs.aws.amazon.com/codepipeline/latest/userguide/detect-state-changes-cloudwatch-events.html
// # https://docs.aws.amazon.com/codebuild/latest/APIReference/API_BuildPhase.html

func BuildAndSendSlackMessage(detail events.CodePipelineEventDetail, buildID string) {
	fmt.Println("## BuildAndSendSlackMessage:")
	var blocks []slack.Block // Full Set of blocks that make up message to be send to slack.

	pipelineTitleText := fmt.Sprintf("*%s*", detail.Pipeline)
	pipelineTitleTextBlock := slack.NewTextBlockObject("mrkdwn", pipelineTitleText, false, false)

	blocks = append(blocks, slack.NewSectionBlock(pipelineTitleTextBlock, nil, nil, slack.SectionBlockOptionBlockID("pipeline_title_text_block")))

	var revisionTextBlocks []*slack.TextBlockObject
	// TODO: Need to not look this up every time
	revisionSummaryTextBlock := RevisionTextBlockBuilder(&revisionTextBlocks, buildID, detail.Pipeline)

	blocks = append(blocks, slack.NewSectionBlock(revisionSummaryTextBlock, revisionTextBlocks, nil, slack.SectionBlockOptionBlockID("revision_details")))

	statusTitleText := slack.NewTextBlockObject("mrkdwn", "*Pipeline Status:*", false, false)

	blocks = append(blocks, slack.NewSectionBlock(statusTitleText, nil, nil, slack.SectionBlockOptionBlockID("status_title_text")))

	var detailTextBlocks []*slack.TextBlockObject
	DetailTextBlockBuilder(&detailTextBlocks, detail)

	blocks = append(blocks, slack.NewSectionBlock(nil, detailTextBlocks, nil, slack.SectionBlockOptionBlockID("status_details")))

	channelID := os.Getenv("CodePipelineSlackChannelID")

	message.CreateOrUpdateMessage(channelID, buildID, blocks, slack.Attachment{})

	// TODO: Log links etc.
}

func RevisionTextBlockBuilder(revisionTextBlocks *[]*slack.TextBlockObject, buildID string, pipelineName string) *slack.TextBlockObject{
	fmt.Println("Message Lookup")
	// Create the session that the CodePipeline service will use.
	awsSession := session.Must(session.NewSession())

	// Create the CodePipeline service client to make the query request with.
	svc := codepipeline.New(awsSession)

	// https://docs.aws.amazon.com/sdk-for-go/api/service/codepipeline/#CodePipeline.GetPipelineExecutionRequest
	req, resp := svc.GetPipelineExecutionRequest(
		&codepipeline.GetPipelineExecutionInput{
			PipelineExecutionId: aws.String(buildID),
			PipelineName:        aws.String(pipelineName),
		})

	err := req.Send()
	// TODO: Clean this up
	if err != nil {
		fmt.Println(err)
	}

	revisionSummaryTextBlock := slack.NewTextBlockObject("mrkdwn", "",false, false)
	revisionURL := ""

	// TODO: Clean this up
	if len(resp.PipelineExecution.ArtifactRevisions) > 0 {
		revisionInfo := resp.PipelineExecution.ArtifactRevisions[0]
		revisionSummaryTextBlock.Text = *revisionInfo.RevisionSummary
		revisionURL = fmt.Sprintf("<%s|Commit: %.7s>", *revisionInfo.RevisionUrl, *revisionInfo.RevisionId)
	}

	// TODO: Clean this up
	revisionTextBlock2 := slack.NewTextBlockObject("mrkdwn", revisionURL,false, false)
	*revisionTextBlocks = append(*revisionTextBlocks, revisionTextBlock2)

	return revisionSummaryTextBlock
}

func DetailTextBlockBuilder(detailTextBlocks *[]*slack.TextBlockObject, detail events.CodePipelineEventDetail ) {
	statusIconMapping := map[string]string{
		"": 									   message.BuildPhasesUnknown,
		string(events.CodePipelineStateStarted):   message.BuildPhasesInProgress,
		events.CodePipelineStateSucceeded:         message.BuildPhasesSucceeded,
		events.CodePipelineStateResumed:           message.BuildPhasesInProgress,
		events.CodePipelineStateFailed:            message.BuildPhasesFailed,
		events.CodePipelineStateCanceled:          message.BuildPhasesStopped,
		events.CodePipelineStateSuperseded:        message.BuildPhasesStopped,
	}

	detailTextBlock := slack.NewTextBlockObject(
			"plain_text",
		    fmt.Sprintf("%s %s", string(statusIconMapping[string(detail.State)]), string(detail.Stage)),
			true,
			false,
	)
	*detailTextBlocks = append(*detailTextBlocks, detailTextBlock)
}