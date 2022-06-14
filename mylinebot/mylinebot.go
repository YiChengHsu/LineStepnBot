package mylinebot

import (
	"line-stepn-bot/config"
	"line-stepn-bot/log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var myLineBot *linebot.Client

func Init() *linebot.Client {

	log.Info(config.Global)

	myLineBot, err := linebot.New(
		config.Global.LineBot.ChannelSecret,
		config.Global.LineBot.ChannelToken,
	)

	if err != nil {
		log.Fatal(err)
	}
	return myLineBot
}
