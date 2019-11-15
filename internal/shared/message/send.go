package message

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

// Initialise Slack API with the Bot Token
//
var api = slack.New(os.Getenv("OAUTH_ACCESS_TOKEN"))

func CreateOrUpdateMessage(channelID string, buildID string, blocks []slack.Block, attachment slack.Attachment) {
	slackTS := ""
	slackTS = SlackTSLookup(buildID)

	if slackTS == "" {
		fmt.Printf("\n%s %s", "Not Found on first attempt:", buildID)
		respTimestamp := CreateMessage(channelID, blocks, attachment)
		SaveNewMessageTS(buildID, respTimestamp)
	} else {
		fmt.Printf("\n%s %s", "Found:", buildID)
		_, _, _, err := api.UpdateMessage(channelID, slackTS, slack.MsgOptionBlocks(blocks...), slack.MsgOptionAttachments(attachment))
		HandleSlackErrors(err, blocks)
	}
}

func CreateMessage(channelID string, blocks []slack.Block, attachment slack.Attachment) string {
	_, respTimestamp, err := api.PostMessage(channelID, slack.MsgOptionBlocks(blocks...), slack.MsgOptionAttachments(attachment))
	HandleSlackErrors(err, blocks)
	fmt.Println("## Saved new Message:")
	fmt.Println(respTimestamp)
	return respTimestamp
}

// Generic Error printer if exists - not much else we can do with them.
//
func HandleSlackErrors(err error, blocks []slack.Block) {
	if err != nil {
		fmt.Println("## Error:")
		fmt.Println(err.Error())
		fmt.Println("## Request:")
		prettyJson, err := json.MarshalIndent(blocks, "", "    ")
		if err != nil {
			fmt.Println("## JSON Indent Error:")
			fmt.Println(err.Error())
		}
		fmt.Printf("\n%s", prettyJson)
	}
}
