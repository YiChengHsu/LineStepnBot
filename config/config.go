package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"line-stepn-bot/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	TimeZone          *time.Location
	AllowOrigins      []string
	HTTPListenAddress string
	HTTPListenPort    int64
	Port              string
	BlockChain        BlockChain
	LineBot           LineBot
}

type BlockChain struct {
	Currencies []string
	ApiUrl     string
}

type LineBot struct {
	ChannelSecret string
	ChannelToken  string
	AlertAccounts []string
}

var Global *Config

func InitConfig() (err error) {
	Global = new(Config)

	// load time zone
	timeZone := os.Getenv("TIME_ZONE")
	if err = validation.Validate(&timeZone, validation.Required); err != nil {
		err = fmt.Errorf(`"TIME_ZONE" %w`, err)
		return
	}
	if Global.TimeZone, err = time.LoadLocation(timeZone); err != nil {
		err = fmt.Errorf(`error on parsing "TIME_ZONE": %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "TIME_ZONE", timeZone))

	// load http address
	Global.HTTPListenAddress = os.Getenv("HTTP_LISTEN_ADDR")
	if err = validation.Validate(&Global.HTTPListenAddress, validation.Required, is.IPv4); err != nil {
		err = fmt.Errorf(`"HTTP_LISTEN_ADDR %w"`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "HTTP_LISTEN_ADDR", Global.HTTPListenAddress))

	// // load http port
	// httpListenPortString := os.Getenv("HTTP_LISTEN_PORT")
	// if err = validation.Validate(&httpListenPortString, validation.Required, is.Int); err != nil {
	// 	err = fmt.Errorf(`"HTTP_LISTEN_PORT %w`, err)
	// 	return
	// }
	// if Global.HTTPListenPort, err = strconv.ParseInt(httpListenPortString, 10, 64); err != nil {
	// 	err = fmt.Errorf(`error on parsing "HTTP_LISTEN_PORT": %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "HTTP_LISTEN_PORT", httpListenPortString))

	// load http port on heroku
	Global.Port = os.Getenv("PORT")
	if err = validation.Validate(&Global.Port, validation.Required, is.Int); err != nil {
		err = fmt.Errorf(`"PORT %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "PORT", Global.Port))

	// load allow origins
	allowOriginStrings := os.Getenv("ALLOW_ORIGINS")
	Global.AllowOrigins = strings.Split(allowOriginStrings, ",")
	if err = validation.Validate(Global.AllowOrigins, validation.Required, validation.Each(is.URL)); err != nil {
		err = fmt.Errorf(`"ALLOW_ORIGINS" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "ALLOW_ORIGINS", allowOriginStrings))

	// load currencies
	currencyStrings := os.Getenv("BLOCKCHAIN_CURRENCY")
	Global.BlockChain.Currencies = strings.Split(currencyStrings, ",")
	if err = validation.Validate(Global.BlockChain.Currencies, validation.Required); err != nil {
		err = fmt.Errorf(`"BLOCKCHAIN_CURRENCY" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "BLOCKCHAIN_CURRENCY", currencyStrings))

	// load api url
	Global.BlockChain.ApiUrl = os.Getenv("BLOCKCHAIN_API_URL")
	if err = validation.Validate(&Global.BlockChain.ApiUrl, validation.Required, is.URL); err != nil {
		err = fmt.Errorf(`"BLOCKCHAIN_API_URL" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "BLOCKCHAIN_API_URL", Global.BlockChain.ApiUrl))

	// load linebot secret
	Global.LineBot.ChannelSecret = os.Getenv("LINE_CHANNEL_SECRET")
	if err = validation.Validate(&Global.LineBot.ChannelSecret, validation.Required); err != nil {
		err = fmt.Errorf(`"LINE_CHANNEL_SECRET" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "LINE_CHANNEL_SECRET", Global.LineBot.ChannelSecret))

	// load linebot token
	Global.LineBot.ChannelToken = os.Getenv("LINE_CHANNEL_TOKEN")
	if err = validation.Validate(&Global.LineBot.ChannelToken, validation.Required); err != nil {
		err = fmt.Errorf(`"LINE_CHANNEL_TOKEN" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "LINE_CHANNEL_TOKEN", Global.LineBot.ChannelToken))

	// load alert line account
	accountStrings := os.Getenv("LINE_ALERT_ACCOUNTS")
	Global.LineBot.AlertAccounts = strings.Split(accountStrings, ",")
	if err = validation.Validate(Global.LineBot.AlertAccounts, validation.Required); err != nil {
		err = fmt.Errorf(`"LINE_ALERT_ACCOUNTS" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "LINE_ALERT_ACCOUNTS", accountStrings))

	return
}
