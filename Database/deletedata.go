package database

import (
	"strconv"

	logging "mikel-kunze.com/uploadservice/Logging"
)

// Deletes a user by the given id and returns a bool to indicate succes
func DeleteUser(userID uint) bool {
	dbcon := CreateDBCon()

	if dbcon == nil {
		logging.LogEntry("[Error]", "Cannot connect to database")
		return false
	}

	defer dbcon.Close()

	if _, err := dbcon.Exec("DELETE FROM Users WHERE id = ?", userID); err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}

	msg := "Deleted user " + strconv.FormatUint(uint64(userID), 10)

	logging.LogEntry("[Information]", msg)

	return true
}

// Deletes a user directory by the given UserDirectorys struct and returns a bool to indicate succes
func DeleteUserDir(dir *UserDirectorys) bool {
	dbcon := CreateDBCon()

	if dbcon == nil {
		logging.LogEntry("[Error]", "Cannot connect to database")
		return false
	}

	defer dbcon.Close()

	if _, err := dbcon.Exec("DELETE FROM UserDirectorys WHERE DirID = ?", dir.DirID); err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}

	msg := "Deleted following dir:" + dir.DirName

	logging.LogEntry("[Error]", msg)

	return true
}

// TODO: implement func to delete multiple dirs
func DeleteUserDirs(dirs *[]UserDirectorys) {

}

// Deletes a access token by the given struct
func DeleteAccesstoken(token ActiveAccessTokens) {

}
