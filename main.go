package main

import (
	"line-stepn-bot/config"
	"line-stepn-bot/cron"
	"line-stepn-bot/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.ErrorLogger())

	config.InitConfig()

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
