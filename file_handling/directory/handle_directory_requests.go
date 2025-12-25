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

	authorized, claims := authentication.AuthorizeWithToken(r.Header.Get("Authorization"))

	if !authorized {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	switch r.Method {
	// If get --> should return all User dirs
	case "GET":
		userDirs := database.GetUserDirs(claims.UserID)

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

		body, err := io.ReadAll(r.Body)

		if err != nil {
			logging.LogEntry("[Error]", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var userDir database.UserDirectorys

		if err := json.Unmarshal(body, &userDir); err != nil {
			logging.LogEntry("[Error]", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		success := CreateUserDir(userDir)

		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			break
		}
	// If delete --> delete the given dir
	case "DELETE":
		success := DeleteUserDir(database.GetDirectoryByName(claims.UserDirectory))
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
