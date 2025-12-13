package files

import (
	"net/http"

	authentication "mikel-kunze.com/uploadservice/Authentication"
	logging "mikel-kunze.com/uploadservice/Logging"
)

// TODO: check function and refactor
func HttpFileUploadRequest(w http.ResponseWriter, r *http.Request) {

	// checks if the client is autorized
	authorized, userName := authentication.AuthorizeWithToken(r.Header.Get("Authorization"))

	if !authorized || userName == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	logging.LogEntry("[Access]: Upload ", authentication.GetIP(r))
	r.ParseMultipartForm(20 << 30)

	switch r.Method {
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
	}

	w.WriteHeader(http.StatusOK)
}
