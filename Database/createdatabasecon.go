package database

import (
	"database/sql"
	"os"

	logging "mikel-kunze.com/uploadservice/Logging"
)

func CreateDBCon() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("connection-string"))

	if err != nil {
		logging.LogEntry("[Error]", err.Error())
		return nil
	}
	return db
}
