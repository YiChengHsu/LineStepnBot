package main

import (
	"line-stepn-bot/config"
	"line-stepn-bot/cron"
	"line-stepn-bot/handler"
	"line-stepn-bot/log"
	"line-stepn-bot/mylinebot"

	"github.com/gin-gonic/gin"
)

func main() {

	err := config.InitConfig()
	if err != nil {
		log.Fatal(log.LabelStartup, "Failed to start. ", err)
	}

	mylinebot.Init()
	err = mylinebot.Init()
	if err != nil {
		log.Fatal(log.LabelStartup, "Failed to start. ", err)
	}

	r := gin.Default()
	r.Use(gin.ErrorLogger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/coins", handler.CurrencyHandler)
	r.POST("/callback", handler.LineHandler())

	cron.InitCron()

	// r.Run(fmt.Sprintf(":%s", config.Global.Port))
	r.Run()
}
