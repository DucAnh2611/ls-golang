package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Logger *log.Logger
)

func InitLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}

	Logger = log.New(file, "", log.LstdFlags|log.Lshortfile)

	multi := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	Logger.SetOutput(multi.Writer())
}

func Info(message string) {
	Logger.SetPrefix("[INFO] ")
	Logger.Println(message)
}

func Error(message string) {
	Logger.SetPrefix("[ERROR] ")
	Logger.Println(message)
}

func Debug(message string) {
	Logger.SetPrefix("[DEBUG] ")
	Logger.Println(message)
}

func logTimestamp() string {
	return time.Now().Format(time.RFC3339)
}
