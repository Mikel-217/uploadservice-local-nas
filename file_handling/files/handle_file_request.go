package files

import (
	"encoding/json"
	"net/http"

	authentication "mikel-kunze.com/uploadservice/authentication"
	"mikel-kunze.com/uploadservice/database"
	logging "mikel-kunze.com/uploadservice/logging"
)

// HTTP-Header field with the directoryID -> maybe other solution? -> base64 encoded

// TODO: check function and refactor
func HttpFileUploadRequest(w http.ResponseWriter, r *http.Request) {

	// checks if the client is autorized
	// Claims are currently not needed
	authorized, _ := authentication.AuthorizeWithToken(r.Header.Get("Authorization"))

	if !authorized {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	logging.LogEntry("[Access]: Upload ", authentication.GetIP(r))
	r.ParseMultipartForm(20 << 30)

	switch r.Method {
	case "GET":

		// gets all files from the given user
		files := database.GetUserFiles(0)

		data, err := json.Marshal(files)

		if err != nil {
			logging.LogEntry("[Error]", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	case "POST":
		if err := HandleUpload(*r.MultipartForm, r.Header.Get("Authorization")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "DELETE":
		if err := DeleteFiles(*r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
