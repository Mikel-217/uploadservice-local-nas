package startup

import (
	"database/sql"
	"errors"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"mikel-kunze.com/uploadservice/logging"
)

// does some things on startup
func OnServerStartup() (bool, error) {
	succesDB := checkDB()

	if succesDB {
		return true, nil
	} else {
		errorMssg := "failed to connect to database:" + strconv.FormatBool(succesDB)
		return false, errors.New(errorMssg)
	}
}

// checks if the Database is online
// TODO: dont forgett to add connection string to enviroment variabless
func checkDB() bool {
	if _, err := sql.Open("mysql", os.Getenv("connection-string")); err != nil {
		logging.LogEntry("[Error]", err.Error())
		return false
	}
	return true
}

// gets the JWT-Key out of enviorment variables --> Key is the jwt secret
func GetKey() []byte {
	return []byte(os.Getenv("jwt-key"))
}
