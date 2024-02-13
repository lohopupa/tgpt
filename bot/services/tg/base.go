package tg

import (
	"bot/types"
	"bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


type Bot struct {
	bot *tgbotapi.BotAPI
	botType types.BotType
}

type BaseHandler interface {
	Init(config config.Config) (Bot, error)
	Start() (error)
	Stop()
}