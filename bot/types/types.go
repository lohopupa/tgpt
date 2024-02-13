package types

type BotType int

const (
	BotTypeLogPoll BotType = iota
	BotTypeWebHook
)