package config

import (
	logger "bot/common/logger"
	types "bot/types"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	LOG_LEVEL logger.LOG_LEVEL

	TG_API_KEY string
	TG_WH_URL  string
	TG_WH_PORT int
	TG_TYPE    types.BotType
}

func GetConfig() Config {
	return Config{
		LOG_LEVEL:  getEnvLoggingLevel("LOG_LEVEL", logger.ERROR),
		TG_API_KEY: getEnvStr("TG_API_KEY", ""),
		TG_WH_URL:  getEnvStr("TG_WH_URL", "http://localhost"),
		TG_WH_PORT: getEnvInt("TG_WH_PORT", 8080),
		TG_TYPE:    getEnvBotType("TG_BOT_TYPE", types.BotTypeLogPoll),
	}
}

func getEnvStr(name string, defaultValue string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return defaultValue
}
func getEnvInt(name string, defaultValue int) int {
	if v := os.Getenv(name); v != "" {
		if int_v, err := strconv.Atoi(v); err == nil {
			return int_v
		}
	}
	return defaultValue
}
func getEnvBool(name string, defaultValue bool) bool {
	if v := os.Getenv(name); v != "" {
		return true
	}
	return false
}
func getEnvLoggingLevel(name string, defaultValue logger.LOG_LEVEL) logger.LOG_LEVEL {
	if v := os.Getenv(name); v != "" {
		v = strings.ToUpper(v)
		switch v {
		case "INFO":
			return logger.INFO
		case "WARNING":
			return logger.WARNING
		case "ERROR":
			return logger.ERROR
		case "DEBUG":
			return logger.DEBUG
		}
	}
	return defaultValue
}
func getEnvBotType(name string, defaultValue types.BotType) types.BotType {
	if v := os.Getenv(name); v != "" {
		switch v {
		case "WH":
			return types.BotTypeLogPoll
		case "LP":
			return types.BotTypeWebHook
		}
	}
	return defaultValue
}
