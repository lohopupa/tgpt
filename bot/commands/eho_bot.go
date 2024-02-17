package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateEchoBotCommands() []CommandHandler {
	handlers := make([]CommandHandler, 1)
	handlers[0] = CommandHandler{
		RegExp: `.*`,
		Handler: func(update *tgbotapi.Update) ([]tgbotapi.Chattable, error) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("```\n%s\n```", update.Message.Text))
			msg.ParseMode = "MarkdownV2"
			msgs := []tgbotapi.Chattable{msg}
			return msgs, nil
		},
	}
	return handlers
}
