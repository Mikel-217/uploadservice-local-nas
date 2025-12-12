package main

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"net/http"

	database "mikel-kunze.com/uploadservice/Database"
	logging "mikel-kunze.com/uploadservice/Logging"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil || string(body) == "" {
		logging.LogEntry("[Error]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user database.UserStruct

	if err := json.Unmarshal(body, &user); err != nil {
		logging.LogEntry("[Error]", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	success := database.CreateNewUser(user)

	if !success {
		logging.LogEntry("[Error]", "Failed to create User")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		logging.LogEntry("[Error]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	succes := database.DeleteUser(uint(binary.BigEndian.Uint64(body)))

	if !succes {
		logging.LogEntry("[Error]", "Failed to delete user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
