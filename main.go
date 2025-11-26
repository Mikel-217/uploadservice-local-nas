package main

import (
	"net/http"

	authen "mikel-kunze.com/uploadservice/Authentication"
	logging "mikel-kunze.com/uploadservice/Logging"
	upload "mikel-kunze.com/uploadservice/Upload"
)

// TODO:
// - Add authentication

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth", sendNewAccess)
	mux.HandleFunc("/api/file/2/", httpFileUploadRequest)

	http.ListenAndServe(":8080", mux)
}

func httpFileUploadRequest(w http.ResponseWriter, r *http.Request) {

	// checks if the client is autorized
	authorized := authen.Authenticate(w.Header().Get("Authorization"))

	if !authorized {
		http.Error(w, "forbitten", http.StatusForbidden)
		return
	}

	logging.LogEntry("[Access]: Upload ", authen.GetIP(r))
	r.ParseMultipartForm(20 << 30)

	query := r.URL.Query()
	homeDirUser := query.Get("user-homedir")

	if homeDirUser == "" {
		http.Error(w, "User is null", http.StatusBadRequest)
		return
	}

	if err := upload.HandleUpload(*r.MultipartForm, homeDirUser); err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendNewAccess(w http.ResponseWriter, r *http.Request) {

	http.Error(w, "Not implemented", http.StatusNotFound)
}
