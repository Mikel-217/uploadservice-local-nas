package database

import (
	"strconv"

	logging "mikel-kunze.com/uploadservice/logging"
)

// Deletes a user by the given id and returns a bool to indicate success
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

// Deletes a user directory by the given UserDirectorys struct and returns a bool to indicate success
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

// Deletes directory by the given slice of the UserDirectorys struct.
// Returns a bool to indicate success
func DeleteUserDirs(dirs []UserDirectorys) bool {
	dbcon := CreateDBCon()

	if dbcon == nil {
		logging.LogEntry("[Error]", "Cannot connect to database")
		return false
	}

	defer dbcon.Close()

	for _, dir := range dirs {
		if _, err := dbcon.Exec("", dir.DirID); err != nil {
			logging.LogEntry("[Error]", err.Error())
			continue
		}

		msg := "Deleted dir:" + dir.DirPath
		logging.LogEntry("[Information]", msg)
	}

	return true
}

// Deletes a access token by the given struct
func DeleteAccesstoken(token *ActiveAccessTokens) bool {
	dbcon := CreateDBCon()

	if dbcon == nil {
		logging.LogEntry("[Error]", "Cannot connect to database")
		return false
	}

	defer dbcon.Close()

	if _, err := dbcon.Exec("DELETE FROM ActiveAccessTokens WHERE TokenID = ?", token.TokenID); err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}

	return true
}

// Deletes a file in the database by the given struct
func DeleteUserFile(file *UserFiles) bool {

	return true
}
