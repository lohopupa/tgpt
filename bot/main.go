package main

import (
	logger "bot/common/logger"
	"bot/config"
)

func main(){
	appConf := config.GetConfig()
	logger.SetLogLevel(appConf.LOG_LEVEL)
	// for {
		logger.Info("Bot Started")
	// }
	
}