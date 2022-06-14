package handler

import (
	"fmt"
	"net/http"
	"strings"

	"line-stepn-bot/currency"
	"line-stepn-bot/log"
	"line-stepn-bot/mylinebot"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var getMessage = "!s"

func LineHandler() gin.HandlerFunc {

	var myBot = mylinebot.Init()

	return func(c *gin.Context) {
		events, err := myBot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.JSON(http.StatusBadRequest, nil)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:

					log.Info(message)

					if message.Text == getMessage {
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							NewEmojiMsg(message),
						).Do(); err != nil {
							log.Error(err)
						}
					}
				case *linebot.ImageMessage:
					log.Info(message)
				case *linebot.VideoMessage:
					log.Info(message)
				case *linebot.AudioMessage:
					log.Info(message)
				case *linebot.FileMessage:
					log.Info(message)
				case *linebot.LocationMessage:
					log.Info(message)
				case *linebot.StickerMessage:
					log.Info(message)
				default:
					log.Info("Unknown message: %v", message)
				}
			default:
				log.Info("Unknown event: %v", event)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"success": events,
		})

	}
}

func CurrencyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": currency.CurrencyData,
	})
}

func NewEmojiMsg(msg *linebot.TextMessage) linebot.SendingMessage {

	emojiIndex := 0
	emojiProductId := "5ac21ef5031a6752fb806d5e"

	repMsg := linebot.NewTextMessage("現在幣價資訊:\n\n")
	totalMsg := &repMsg.Text

	for _, currency := range currency.CurrencyData {

		data := currency.MarketData

		singleText := fmt.Sprintf(
			"%s:\n價格:  %.2f $ \n 24小時漲跌:  %.2f\n\n",
			strings.ToUpper(currency.Symbol),
			data.CurrencyPrice.USD,
			data.PriceChangePercentage24H.USD,
		)

		var emojiId string
		switch change := data.PriceChangePercentage24H.USD; {
		case change > 0:
			emojiId = "050"
		case change <= 0:
			emojiId = "037"
		}

		emoji := linebot.Emoji{
			Index:     emojiIndex,
			ProductID: emojiProductId,
			EmojiID:   emojiId,
		}

		*totalMsg = fmt.Sprint(repMsg, singleText)
		repMsg.AddEmoji(&emoji)

		emojiIndex++
	}

	return repMsg
}
