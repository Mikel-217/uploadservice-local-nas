package directory

import (
	"encoding/json"
	"io"
	"net/http"

	database "mikel-kunze.com/uploadservice/Database"
)

// TODO: add authentication!!
func HttpDirRequest(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userDir database.UserDirectorys

	if err := json.Unmarshal(body, &userDir); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		success := CreateUserDir(userDir)
		if !success {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			break
		}
	case "DELETE":
		success := DeleteUserDir(userDir)
		if !success {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			break
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
