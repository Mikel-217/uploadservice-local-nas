package files

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// TODO: change to dynamic for win and linux
// TODO: refactor
var currFileDir = "mnt/voldata/"
var UserDirection = ""

func HandleUpload(files multipart.Form, authToken string) error {

	if len(files.File["attachments"]) <= 0 {
		return errors.New("Files Empty")
	}

	getUserDir(authToken)

	if _, err := os.Stat(UserDirection); err != nil {
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

// TODO: get user dir via DB
// Gets the user dir
func getUserDir(token string) {
	strings.Replace(token, "Baerer", "", 0)

	// TODO: get data out of jwt and get the user by id
	// Also think about getting the right dir :)

}
