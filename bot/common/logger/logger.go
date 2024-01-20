package common

import (
	"fmt"
	"log"
)

type LOG_LEVEL int

const (
	ERROR LOG_LEVEL = iota
	WARNING
	INFO
	DEBUG
)

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

type Logger struct {
	logLevel LOG_LEVEL
}

func NewLogger() Logger {
	return Logger{
		logLevel: INFO,
	}
}

func (l *Logger) SetLogLevel(logLevel LOG_LEVEL) {
	l.logLevel = logLevel
}

func (l *Logger) log(logLevel LOG_LEVEL, prms ...any) {
	if l.logLevel <= logLevel {
		log.Println(fmt.Sprintf("[%s]: ", logLevel.repr()), prms)
	}
}

func (l *Logger) Info(prms ...any) {
	l.log(INFO, prms...)
}

func (l *Logger) Warn(prms ...any) {
	l.log(WARNING, prms...)
}

func (l *Logger) Err(prms ...any) {
	l.log(ERROR, prms...)
}

func (l *Logger) Panic(prms ...any) {
	log.Panic(prms...)
}

func (l *Logger) Fatal(prms ...any) {
	log.Fatal(prms...)
}
