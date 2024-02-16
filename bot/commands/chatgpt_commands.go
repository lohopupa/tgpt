package commands

import (
	openai "bot/services/openai"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateChatGPTCommands(openaiClient openai.OpenAiClient) []CommandHandler {
	commands := make([]CommandHandler, 0)
	chatQuery := CommandHandler{
		RegExp: ".*",
		Handler: func(update *tgbotapi.Update) (tgbotapi.Chattable, error) {
			history := make([]string, 0)
			resp, err := openaiClient.Query(update.Message.Text, history)
			if err != nil {
				return nil, err
			}
			return tgbotapi.NewMessage(update.Message.Chat.ID, resp), nil
		},
	}
	commands = append(commands, chatQuery)
	return commands

}
