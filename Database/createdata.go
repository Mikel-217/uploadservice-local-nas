package database

import (
	"encoding/base64"
	"time"

	logging "mikel-kunze.com/uploadservice/Logging"
)

// Sets a new Active token to DB
func CreateNewToken(token string, expiration time.Time) {
	db := CreateDBCon()

	if db == nil {
		logging.LogEntry("[Error]", "Cannot connect to db!")
	}

	defer db.Close()

	if _, err := db.Exec("INSERT INTO ActiveAccessTokens (TokenID, ActiveToken, ExpirationDate) VALUES (DEFAULT, ?, ?)", token, expiration); err != nil {
		logging.LogEntry("[Error]", err.Error())
	}
}

// Creates a new user
func CreateNewUser(user UserStruct) {
	db := CreateDBCon()

	if db == nil {
		logging.LogEntry("[Error]", "Cannot connect to db!")
	}

	defer db.Close()

	_, err := db.Exec("INSERT INTO Users (UserID, UserName, UserPW) VALUES (DEFAULT, ?, ?)", user.UserName, base64.StdEncoding.EncodeToString([]byte(user.PW)))

	if err != nil {
		logging.LogEntry("[Error]", err.Error())
	}
}
