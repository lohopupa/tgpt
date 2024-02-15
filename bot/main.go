package main

import (
	logger "bot/common/logger"
	"bot/config"
	"bot/services/tg"
)

func main(){
	config := config.GetConfig()
	logger.SetLogLevel(config.AppConfig.LogLevel)
	bot, err := tg.CreateBot(config.TgConfig)
	if err != nil {
		logger.Fatal(err)
	}
	bot.Start()
}