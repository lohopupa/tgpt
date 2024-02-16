package tg

import (
	"bot/commands"
	"bot/common/logger"
	"bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type LongPollBot struct {
	instance       *tgbotapi.BotAPI
	updateTimeout int
}

func CreateLP(config config.TgConfig) (*LongPollBot, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, err
	}

	// if config. == logger.DEBUG {
	// 	bot.Debug = true
	// }

	logger.Info("Authorized on account %s", bot.Self.UserName)

	x := LongPollBot{
		instance:       bot,
		updateTimeout: config.LongPollUpdateTimeout,
	}
	logger.Info("LongPoll bot created!")
	return &x, nil
}

func (bot LongPollBot) Start(commands commands.Commands) error {
	logger.Info("LongPoll bot started!")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = bot.updateTimeout
	updates := bot.instance.GetUpdatesChan(u)
	commands.Start(bot.instance, updates)

	

	return nil
}

func (bot LongPollBot) Stop() error {
	logger.Info("LongPoll bot stopped!")
	return nil
}
