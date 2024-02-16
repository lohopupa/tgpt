package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateEchoBotCommands() []CommandHandler {
	handlers := make([]CommandHandler, 1)
	handlers[0] = CommandHandler{
		RegExp: `.*`,
		Handler: func(update *tgbotapi.Update) (tgbotapi.Chattable, error) {
			return tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text), nil
		},
	}
	return handlers
}
