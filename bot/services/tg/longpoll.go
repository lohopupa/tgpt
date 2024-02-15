package tg

import (
	logger "bot/common/logger"
	"bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type LongPollBot struct {
	instance       *tgbotapi.BotAPI
	update_timeout int
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
		update_timeout: config.LongPollUpdateTimeout,
	}
	logger.Info("LongPoll bot created!")
	return &x, nil
}

func (bot LongPollBot) Start() error {
	logger.Info("LongPoll bot started!")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = bot.update_timeout
	updates := bot.instance.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			logger.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.instance.Send(msg)
		}
	}

	return nil
}

func (bot LongPollBot) Stop() error {
	logger.Info("LongPoll bot stopped!")
	return nil
}
