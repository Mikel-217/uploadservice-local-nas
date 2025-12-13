package directory

import (
	"errors"
	"os"

	database "mikel-kunze.com/uploadservice/Database"
	logging "mikel-kunze.com/uploadservice/Logging"
)

// Deletes a directory from the User
func DeleteUserDir(dir database.UserDirectorys) bool {

	if _, err := os.Stat(dir.DirPath); errors.Is(err, os.ErrNotExist) {
		database.DeleteUserDir(&dir)
		return true
	} else if err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}

	if err := os.Remove(dir.DirPath); err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}

	return database.DeleteUserDir(&dir)
}
