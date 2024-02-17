package commands

import (
	openai "bot/services/openai"
	// "fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateChatGPTCommands(openaiClient openai.OpenAiClient) []CommandHandler {
	commands := make([]CommandHandler, 0)
	chatQuery := CommandHandler{
		RegExp: ".*",
		Handler: func(update *tgbotapi.Update) ([]tgbotapi.Chattable, error) {
			history := make([]string, 0)
			resp, err := openaiClient.Query(update.Message.Text, history)
			if err != nil {
				return nil, err
			}
			msgs := make([]tgbotapi.Chattable, 0)
			// for idx, content := range strings.Split(resp, "```") {
			// 	var msg tgbotapi.MessageConfig
			// 	if idx%2 == 1 {
			// 		msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("```\n%s\n```", content))
			// 		msg.ParseMode = "MarkdownV2"
			// 	} else {
			// 		msg = tgbotapi.NewMessage(update.Message.Chat.ID, content)
			// 	}
			// 	msgs = append(msgs, msg)
			// }
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, escapeSymbols(resp))
			msg.ParseMode = "MarkdownV2"
			msgs = append(msgs, msg)
			return msgs, nil
		},
	}
	commands = append(commands, chatQuery)
	return commands

}

func escapeSymbols(text string) string {
	lines := strings.Split(text, "\n")
	result := ""
   
	isCodeBlock := false
   
	for _, line := range lines {
	 if strings.HasPrefix(line, "```") {
	   isCodeBlock = !isCodeBlock
	  }
	
	  if !isCodeBlock {
	   symbols := strings.Split(line, "")
	
	   for i, symbol := range symbols {
		if strings.ContainsAny(symbol, "\\*_{}[]()<>#+-.!$|^~=") {
		 symbols[i] = "\\" + symbol
		}
	   }
	
	   line = strings.Join(symbols, "")
	  }
	
	  result += line + "\n"
	 }
	
	 return result
	}
