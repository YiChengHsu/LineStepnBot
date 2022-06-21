package mylinebot

import (
	"fmt"
	"line-stepn-bot/config"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var MyLineBot *linebot.Client

func Init() (err error) {

	MyLineBot, err = linebot.New(
		config.Global.LineBot.ChannelSecret,
		config.Global.LineBot.ChannelToken,
	)

	if err != nil {
		err = fmt.Errorf(`line bot set up error: %w`, err)
	}
	return
}
