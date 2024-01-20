package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotType int

const (
	BotTypeLogPoll BotType = iota
	BotTypeWebHook
)

type Bot struct {
	bot *tgbotapi.BotAPI
	botType BotType
}


func NewBot(token string, botType BotType) {
	
}