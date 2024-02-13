package common

import (
	"log"
)

type LOG_LEVEL int

const (
	ERROR LOG_LEVEL = iota
	WARNING
	INFO
	DEBUG
)

var logLevel LOG_LEVEL

func (ll LOG_LEVEL) repr() string {
	switch ll {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
	default:
		return "UNKNOWN LOG LEVEL"
	}
}

func SetLogLevel(ll LOG_LEVEL) {
	logLevel = ll
}

func customLog(ll LOG_LEVEL, msg any) {
	if ll <= logLevel {
		log.Printf("[%s]: %s\n", ll.repr(), msg)
	}
}

func Info(msg any) {
	customLog(INFO, msg)
}

func Warn(msg any) {
	customLog(WARNING, msg)
}

func Err(msg any) {
	customLog(ERROR, msg)
}

func Panic(msg any) {
	log.Panic(msg)
}

func Fatal(msg any) {
	log.Fatal(msg)
}
