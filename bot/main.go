package main

import (
	"bot/commands"
	"bot/common/logger"
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
	cmds := commands.Commands{
		CmdHandlers: commands.CreateEchoBotCommands(),
	}
	bot.Start(cmds)
}