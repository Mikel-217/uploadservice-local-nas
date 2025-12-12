package files

import (
	"net/http"

	authentication "mikel-kunze.com/uploadservice/Authentication"
	logging "mikel-kunze.com/uploadservice/Logging"
)

// TODO: check function and refactor
func HttpFileUploadRequest(w http.ResponseWriter, r *http.Request) {

	// only accept POST requests
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// checks if the client is autorized
	authorized, userName := authentication.AuthorizeWithToken(r.Header.Get("Authorization"))

	if !authorized || userName == "" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	logging.LogEntry("[Access]: Upload ", authentication.GetIP(r))
	r.ParseMultipartForm(20 << 30)

	if err := HandleUpload(*r.MultipartForm, userName); err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
