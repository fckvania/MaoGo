package helpers

import (
	"io"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type Logger struct{}

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(io.MultiWriter(file, os.Stderr), "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (log Logger) Info(v any) {
	InfoLogger.Println(v)
}

func (log Logger) Warn(v any) {
	WarningLogger.Println(v)
}

func (log Logger) Error(v any) {
	ErrorLogger.Println(v)
}
