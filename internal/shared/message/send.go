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

func CreateOrUpdateMessage(channelID string, buildID string, blocks []slack.Block, attachment slack.Attachment){
	slackTS := SlackTSLookup(buildID)

	if slackTS == "" {
		// TODO: Race condition with DynamoDB - runs to fast?

		_, respTimestamp, err := api.PostMessage(channelID, slack.MsgOptionBlocks(blocks...), slack.MsgOptionAttachments(attachment))
		HandleSlackErrors(err, blocks)
		SaveNewMessageTS(buildID,respTimestamp)
	} else {
		_, _, _, err := api.UpdateMessage(channelID, slackTS, slack.MsgOptionBlocks(blocks...), slack.MsgOptionAttachments(attachment))
		HandleSlackErrors(err, blocks)
	}
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