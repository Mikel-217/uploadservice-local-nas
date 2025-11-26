package logging

import (
	"errors"
	"log"
	"os"
	"time"
)

func GetLogFile() string {
	currDir, _ := os.Getwd()
	currDir += "\\loggs"
	if _, err := os.Stat(currDir); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(currDir, os.ModeDir)
	}

	files, err := os.ReadDir(currDir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileInfo, err := file.Info()

		if err != nil {
			continue
		}

		if time.Time.Equal(time.Now(), fileInfo.ModTime()) {
			return file.Name()
		} else {
			continue
		}
	}

	newLogFile := time.Now().Format("dd-mm-yyyy") + ".txt"

	nFile, err := os.Create(newLogFile)

	if err != nil {
		log.Fatal(err)
	}

	return nFile.Name()
}
