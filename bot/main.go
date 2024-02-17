package main

import (
	commands "bot/commands"
	"bot/common/logger"
	"bot/config"
	"bot/services/tg"
	openai "bot/services/openai"
)

func main(){
	config := config.GetConfig()
	logger.SetLogLevel(config.AppConfig.LogLevel)
	bot, err := tg.CreateBot(config.TgConfig)
	if err != nil {
		logger.Fatal(err)
	}

	openaiClient, err := openai.CreateClient(config.OpenAIConfig)
	if err != nil {
		logger.Fatal(err)
	}

	cmds := commands.Commands{
		CmdHandlers: commands.CreateChatGPTCommands(openaiClient),
		// CmdHandlers: commands.CreateEchoBotCommands(),
	}
	bot.Start(cmds)
}