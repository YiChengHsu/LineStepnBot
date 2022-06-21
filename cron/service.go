package cron

import (
	"line-stepn-bot/currency"
	"line-stepn-bot/handler"
	"line-stepn-bot/log"
	"time"

	"github.com/jasonlvhit/gocron"
)

func InitCron() {
	if err := gocron.Every(10).Seconds().Do(currency.SyncCurrency); err != nil {
		time.Sleep(2 * time.Second)
		log.Fatal(err)
	}

	if err := gocron.Every(1).Hour().Do(handler.DetectHandler); err != nil {
		time.Sleep(2 * time.Second)
		log.Fatal(err)
	}
	gocron.Start()
	currency.SyncCurrency()
	handler.DetectHandler()
}
