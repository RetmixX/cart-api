package log

import (
	"log"
	"os"
)

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func NewInfoLog() *log.Logger {
	return log.New(os.Stdout, "[INFO]\t", log.LstdFlags)
}

func NewErrorLog() *log.Logger {
	return log.New(os.Stdout, "[ERROR]\t", log.LstdFlags|log.Llongfile)
}

func NewLogger() *Logger {
	return &Logger{
		ErrorLog: NewErrorLog(),
		InfoLog:  NewInfoLog(),
	}
}
