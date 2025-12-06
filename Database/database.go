package database

import (
	"time"

	logging "mikel-kunze.com/uploadservice/Logging"
)

// Gets a user by the given username --> Returns a Userstruct
func GetUserByName(userName string) UserStruct {
	db := CreateDBCon()

	if db == nil {
		logging.LogEntry("[Error]", "Cannot connect to db!")
	}

	defer db.Close()

	var user UserStruct

	err := db.QueryRow("SELECT * FROM Users WHERE UserName = ?", userName).Scan(&user.ID, &user.UserName, &user.UserName)

	if err != nil {
		logging.LogEntry("[Error]", err.Error())
	}

	return user
}

// Gets all UserDirs by the given userID --> Returns a slice
func GetUserDirs(userID uint) []UserDirectorys {

	db := CreateDBCon()

	if db == nil {
		logging.LogEntry("[Error]", "Cannot connect to db!")
	}
	userDirSlice := make([]UserDirectorys, 100)

	defer db.Close()

	rows, err := db.Query("SELECT * FROM UserDirectorys WHERE UserID = ?", userID)

	if err != nil {
		logging.LogEntry("[Error]", err.Error())
	}

	for rows.Next() {
		var userDir UserDirectorys

		if err := rows.Scan(&userDir.DirID, &userDir.DirID, &userDir.DirName, &userDir.DirPath); err != nil {
			logging.LogEntry("[Error]", err.Error())
			continue
		}
		userDirSlice = append(userDirSlice, userDir)
	}

	return userDirSlice
}

// Checks if the token is saved in the DB
func CheckTokenExistence(token string) bool {
	db := CreateDBCon()

	if db == nil {
		logging.LogEntry("[Error]", "Cannot connect to db!")
		return false
	}

	defer db.Close()

	var tokenDB ActiveAccessTokens

	if err := db.QueryRow("SELECT * FROM ActiveAccessTokens WHERE ActiveToken = ?", token).Scan(&tokenDB.TokenID, &tokenDB.ExpirationDate, &tokenDB.ActiveToken); err != nil {
		return false
	}

	if tokenDB.ActiveToken == token && tokenDB.ExpirationDate != time.Now() {
		return true
	}

	return true
}
