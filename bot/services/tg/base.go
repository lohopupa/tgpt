package tg

import (
	"bot/commands"
	"bot/config"
	"bot/types"
	"errors"
	"fmt"
)

type BaseHandler interface {
	Start(commands commands.Commands) error
	Stop() error
}

func CreateBot(config config.TgConfig) (BaseHandler, error) {
	switch config.BotType {
	case types.BotTypeLogPoll:
		return CreateLP(config)
	case types.BotTypeWebHook:
		return CreateWH(config)
	default:
		return nil, errors.New(fmt.Sprintf("Could not recognize bot type %d", config.BotType))
	}
}
