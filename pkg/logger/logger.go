package logger

import "log"

type Logger struct {
	LoggerOk  *log.Logger
	LoggerErr *log.Logger
}
