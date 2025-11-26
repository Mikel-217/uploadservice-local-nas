package logging

import (
	"log"
	"os"
	"time"
)

func LogEntry(logType string, content string) {
	file := GetLogFile()

	currLogFile, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	defer currLogFile.Close()

	if _, err := currLogFile.WriteString(time.Now().GoString() + logType + content); err != nil {
		log.Fatal(err)
	}
}
