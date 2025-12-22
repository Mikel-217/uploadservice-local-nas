package files

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	authentication "mikel-kunze.com/uploadservice/authentication"
	database "mikel-kunze.com/uploadservice/database"
	logging "mikel-kunze.com/uploadservice/logging"
)

// TODO: change to dynamic for win and linux
// TODO: refactor
var currFileDir = "mnt/voldata/"
var UserDirection = ""
var TokenClaims = authentication.Claims{}

// Creates files given by the request
func HandleUpload(files multipart.Form, authToken string) error {

	if len(files.File["attachments"]) <= 0 {
		return errors.New("Files Empty")
	}

	if err := getUserDir(authToken); err != nil {
		logging.LogEntry("[Error]", err.Error())
		return err
	}

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

		userFile := database.UserFiles{
			FileName: createtFile.Name(),
			FilePath: path.Join(UserDirection, createtFile.Name()),
			DirID:    database.GetDirectoryByName(TokenClaims.UserDirectory).DirID,
			UserID:   TokenClaims.UserID,
		}

		if !database.CreateNewFile(&userFile) {
			os.Remove(path.Join(UserDirection, createtFile.Name()))
		}
	}
	return nil
}

// Gets the user dir
// The user direction is a claim in the JWT
func getUserDir(token string) error {

	if token == "" {
		logging.LogEntry("[Error]", "UserToken is null")
		return errors.New("UserToken is null")
	}

	strings.Replace(token, "Baerer", "", 0)

	claims := authentication.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { return authentication.JWTKey, nil })

	if err != nil {
		logging.LogEntry("[Error]", err.Error())
		return err
	}

	tokenClaims := tkn.Claims.(*authentication.Claims)

	if tokenClaims.UserDirectory != "" {
		logging.LogEntry("[Error]", "User directory was null")
		return errors.New("User directory was null")
	}

	UserDirection = tokenClaims.UserDirectory
	return nil
}
