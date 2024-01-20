package main

import (
	logger "bot/common/logger"
	"bot/config"
)

func main(){
	appConf := config.GetConfig()
	log := logger.NewLogger()
	log.SetLogLevel(appConf.LOG_LEVEL)
	
}