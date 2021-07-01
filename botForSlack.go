package main

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"io"
	"log"
	"os"
)

type SlackBotChannelsInfo struct {
	BotToken string        `json:"bot_token"`
	Channels []ChannelType `json:"channels"`
}

type ChannelType struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func main() {

	jsonFile, err := os.Open("example.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened example.json")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var slackBotChannelsInfo SlackBotChannelsInfo

	json.Unmarshal(byteValue, &slackBotChannelsInfo)

	OAuthToken := slackBotChannelsInfo.BotToken

	for i := 0; i < len(slackBotChannelsInfo.Channels); i++ {
		channel := slackBotChannelsInfo.Channels[i]
		sendMessageToSlack(channel.Channel, channel.Text, OAuthToken)
	}
}

func sendMessageToSlack(channelName string, message string, OAuthToken string) {

	api := slack.New(OAuthToken)

	attachment := slack.Attachment{
		Text: message,
	}

	channelId, timestamp, text, err := api.SendMessage(
		channelName,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}
	log.Printf("Message successfully sent to Channel %s at %s\n, text: %s", channelId, timestamp, text)
}
