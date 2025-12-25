package directory

import (
	"encoding/json"
	"io"
	"net/http"

	authentication "mikel-kunze.com/uploadservice/authentication"
	database "mikel-kunze.com/uploadservice/database"
	logging "mikel-kunze.com/uploadservice/logging"
)

// Handels the requests to create or delete a directory
func HttpDirRequest(w http.ResponseWriter, r *http.Request) {

	authorized, userName := authentication.AuthorizeWithToken(r.Header.Get("Authorization"))

	// TODO: revisit me!!!!
	// Bad practice to send stuff in a GET-Request Body!!

	if !authorized || userName == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

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
	// If get --> should return all User dirs
	case "GET":
		userDirs := database.GetUserDirs(userDir.UserID)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		data, err := json.Marshal(userDirs)

		if err != nil {
			logging.LogEntry("[Error]", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	// If post --> create a new User dir
	case "POST":
		success := CreateUserDir(userDir)
		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			break
		}
	// If delete --> delete the given dir
	case "DELETE":
		success := DeleteUserDir(userDir)
		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			break
		}
	// If none of those obove -> bad request
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// always write 200 at the end, if nothing else was send --> bad pracitce ?
	w.WriteHeader(http.StatusOK)
}
