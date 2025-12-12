package logging

import (
	"errors"
	"log"
	"os"
	"path"
	"time"
)

func GetLogFile() string {
	currDir, _ := os.Getwd()
	currDir += "\\logs"

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

		if time.Now().Format(time.DateOnly) == fileInfo.ModTime().Format(time.DateOnly) {
			return path.Join(currDir, file.Name())
		} else {
			continue
		}
	}

	date := time.Now().Truncate(24 * time.Hour)
	const formatString = "02-01-2006"
	newLogFile := date.Format(formatString) + ".log"

	nFile, err := os.Create(path.Join(currDir, newLogFile))

	if err != nil {
		log.Fatal(err)
	}

	defer nFile.Close()

	return nFile.Name()
}
