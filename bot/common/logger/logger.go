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

func customLog(ll LOG_LEVEL, fmtString string, prms... any) {
	if ll <= logLevel {
		log.Printf("[%s]: ", ll.repr())
		log.Printf(fmtString, prms...)
	}
}

func Info(msg any) {
	Infof("%s\n", msg)
}

func Warn(msg any) {
	Warnf("%s\n", msg)
}

func Err(msg any) {
	Errf("%s\n", msg)
}

func Infof(fmtString string, prms... any) {
	customLog(INFO, fmtString, prms...)
}

func Warnf(fmtString string, prms... any) {
	customLog(WARNING, fmtString, prms...)
}

func Errf(fmtString string, prms... any) {
	customLog(ERROR, fmtString, prms...)
}

func Panic(prms... any) {
	log.Panic(prms...)
}

func Fatal(prms... any) {
	log.Fatal(prms...)
}
