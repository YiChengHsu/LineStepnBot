package handler

import (
	"fmt"
	"math"
	"net/http"

	"line-stepn-bot/config"
	"line-stepn-bot/currency"
	"line-stepn-bot/log"
	"line-stepn-bot/mylinebot"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var getMessage = "$"

func LineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		myBot := mylinebot.MyLineBot
		events, err := myBot.ParseRequest(c.Request)

		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.JSON(http.StatusBadRequest, err)
			} else {
				c.JSON(http.StatusInternalServerError, err)
			}
		}

		for _, event := range events {
			log.Info(event.Source.UserID)
			log.Info(event.Source.GroupID)
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:

					if message.Text == getMessage {

						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							NewEmojiMsg(),
						).Do(); err != nil {
							log.Error("here", err)
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

			c.JSON(http.StatusOK, gin.H{
				"success": events,
			})
		}

	}
}

func CurrencyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": currency.CurrencyData,
	})
}

func DetectHandler() (err error) {

	log.Info("Start to detect!")

	for _, data := range currency.CurrencyData {

		hourChange := data.MarketData.PriceChangePercentage1H.USD

		if hourChange > 3 || hourChange < -3 {

			template := "🥬🥬🥬 韭菜警報 🥬🥬🥬 \n\n"
			var light, arrow, trend, zora string

			switch {
			case hourChange > 3:
				light = "💚"
				arrow = "📈"
				trend = "漲"
				zora = "我還沒上車啊 💔💔💔"
			case hourChange < 3:
				light = "❤️"
				arrow = "📉"
				trend = "跌"
				zora = "塊陶啊 🏃💨💨💨"
			}

			template = fmt.Sprintf("%s%s %s 一小時內%s了%.2f%%\n目前價格: %.2f $ \n24H漲跌: %.2f %% %s\n\nZora 表示: %s",
				template,
				light,
				data.Name,
				trend,
				math.Abs(hourChange),
				data.MarketData.CurrencyPrice.USD,
				data.MarketData.PriceChangePercentage24H.USD,
				arrow,
				zora,
			)

			for _, member := range config.Global.LineBot.AlertAccounts {
				_, err = mylinebot.MyLineBot.PushMessage(
					member,
					linebot.NewTextMessage(template),
				).Do()

				if err != nil {
					err = fmt.Errorf("ALERT ERROR: %w", err)
					return
				}
			}

		}
	}

	return
}

func NewEmojiMsg() linebot.SendingMessage {

	// productId := "5ac21ef5031a6752fb806d5e"
	// var emojiArr []string

	totalMsg := "🥰 親愛的韭菜\n現在幣價資訊\n\n"

	for _, currency := range currency.CurrencyData {

		data := currency.MarketData

		var arrow string
		var light string
		switch change := data.PriceChangePercentage24H.USD; {
		case change > 0:
			light = "💚"
			arrow = "📈"
		case change <= 0:
			light = "❤️"
			arrow = "📉"
		}

		singleText := fmt.Sprintf(
			"%s %s\n價格:  %.2f $ \n24小時漲跌: %.2f %% %s\n\n",
			light,
			currency.Name,
			data.CurrencyPrice.USD,
			data.PriceChangePercentage24H.USD,
			arrow,
		)

		totalMsg += singleText
	}

	totalMsg += "Zora 關心你的荷包 😘😘😘"

	return linebot.NewTextMessage(totalMsg)
}
