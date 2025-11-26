package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// TODO: change to dynamic
var currFileDir = "mnt/voldata/"
var UserDirection = ""

func HandleUpload(files multipart.Form, userFileDir string) error {

	if len(files.File["attachments"]) <= 0 {
		return errors.New("Files Empty")
	}

	getUserDir(userFileDir)

	_, err := os.Stat(UserDirection)

	if err != nil {
		return err
	}

	for _, file := range files.File["attachments"] {
		createtFile, err := os.Create(path.Join(UserDirection, file.Filename))

		if errors.Is(err, os.ErrExist) {
			os.Remove(path.Join(UserDirection, file.Filename))
			createtFile, _ = os.Create(path.Join(UserDirection, file.Filename))
		} else {
			return err
		}

		defer createtFile.Close()

		uploadetFile, err := file.Open()

		if err != nil {
			return err
		}

		defer uploadetFile.Close()

		if _, err := io.Copy(createtFile, uploadetFile); err != nil {
			return err
		}
	}
	return nil
}

// TODO: get parent path out of config.json

// Gets the user dir
// if username contains an ";" it has specifiet an direction
// userName;Directory
func getUserDir(userFileDir string) {

	if strings.Contains(userFileDir, ";") {
		splitetusr := strings.Split(userFileDir, ";")
		UserDirection = path.Join(currFileDir, splitetusr[0], splitetusr[1])
	}
	UserDirection = path.Join(currFileDir, userFileDir)
}
