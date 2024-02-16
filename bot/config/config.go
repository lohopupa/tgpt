package config

import (
	logger "bot/common/logger"
	types "bot/types"
	"os"
	"strconv"
	"strings"
)

type TgConfig struct {
	ApiKey                string
	BotType               types.BotType
	LongPollUpdateTimeout int
	WebHookAddr            string
	WebHookLocalPort      int
	WebHookCertFolder     string
}

type AppConfig struct {
	LogLevel logger.LOG_LEVEL
}

type OpenAIConfig struct {
	Token string
	BaseURL string 
}

type Config struct {
	AppConfig AppConfig
	TgConfig  TgConfig
	OpenAIConfig OpenAIConfig
}

func GetConfig() Config {
	return Config{
		AppConfig: AppConfig{
			LogLevel: getEnvLoggingLevel("LOG_LEVEL", logger.ERROR),
		},
		TgConfig: TgConfig{
			ApiKey:                getEnvStr("TG_API_KEY", ""),
			BotType:               getEnvBotType("TG_BOT_TYPE", types.BotTypeLogPoll),
			LongPollUpdateTimeout: getEnvInt("TG_LP_UPDATE_TIMEOUT", 60),
			WebHookAddr:            getEnvStr("TG_WH_ADDR", "localhost"),
			WebHookLocalPort:      getEnvInt("TG_WH_LOCAL_PORT", 8443),
			WebHookCertFolder:     getEnvStr("TG_WH_CERT_FOLDER", "."),
		},
		OpenAIConfig: OpenAIConfig{
			BaseURL: getEnvStr("OPENAI_BASE_URL", "https://api.openai.com/v1"),
			Token: getEnvStr("OPENAI_TOKEN", "token"),
		},
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
			return types.BotTypeWebHook
		case "LP":
			return types.BotTypeLogPoll
		}
	}
	return defaultValue
}
