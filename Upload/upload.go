package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

var currFileDir = "mnt/voldata/"
var UserDirection = ""

func HandleUpload(files multipart.Form, userName string) error {

	if len(files.File["attachments"]) <= 0 {
		return errors.New("Files Empty")
	}

	getUserDir(userName)

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

// Gets the user dir
// if username contains an ";" it has specifiet an direction
// userName;Directory
func getUserDir(userName string) {

	if strings.Contains(userName, ";") {
		splitetusr := strings.Split(userName, ";")
		UserDirection = path.Join(currFileDir, splitetusr[0], splitetusr[1])
	}
	UserDirection = path.Join(currFileDir, userName)
}
