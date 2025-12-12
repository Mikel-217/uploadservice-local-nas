package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// FIXMEEEEEE --> only one logentryw
func LogEntry(logType string, content string) {
	file := GetLogFile()

	currLogFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer currLogFile.Close()

	if _, err := currLogFile.WriteString("\n" + time.Now().Format("2006/01/02 15:04:05") + logType + content); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logged entry")
}
