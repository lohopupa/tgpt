package tg

import (
	"bot/commands"
	"bot/config"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebHookBot struct {
	instance *tgbotapi.BotAPI
	config config.TgConfig
}

func CreateWH(config config.TgConfig) (*WebHookBot, error) {
	return nil, errors.New("WebHook API is not supported yet")
}

func (bot WebHookBot) Start(commands commands.Commands) error {
	return errors.New("WebHook API is not supported yet")
}


func (bot WebHookBot) Stop() error {
	return errors.New("WebHook API is not supported yet")
}


