package directory

import (
	"errors"
	"os"

	database "mikel-kunze.com/uploadservice/Database"
	logging "mikel-kunze.com/uploadservice/Logging"
)

func CreateUserDir(dir database.UserDirectorys) bool {

	if _, err := os.Stat(dir.DirPath); errors.Is(err, os.ErrExist) {
		return true
	} else if err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}

	if err := os.Mkdir(dir.DirPath, 0755); err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}

	return database.CreateNewUserDirectory(&dir)
}
