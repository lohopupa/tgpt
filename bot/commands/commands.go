package commands

import (
	"bot/common/logger"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commands struct {
	CmdHandlers []CommandHandler
}

func (this Commands) Start(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil {
			logger.Info("[%s]: %s", update.Message.From.UserName, update.Message.Text)
			for _, cmd := range this.CmdHandlers {
				if matched, _ := regexp.MatchString(cmd.RegExp, update.Message.Text); matched {
					newMsgs, err := cmd.Handler(&update)
					if err != nil {
						logger.Err(err.Error())
						bot.Send(errorMsg(&update))
					}
					for idx, msg := range newMsgs {
						logger.Info("Send %d messge", idx)
						_, err := bot.Send(msg)
						if err != nil {
							logger.Err(err.Error())
							bot.Send(errorMsg(&update))
						}
					}
					break
				}
			}
		}
	}
	return nil
}

type CommandHandler struct {
	RegExp  string
	Handler func(update *tgbotapi.Update) ([]tgbotapi.Chattable, error)
}

func errorMsg(update *tgbotapi.Update) tgbotapi.Chattable {
	return tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong, try  again later!")
}
