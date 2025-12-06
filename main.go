package main

import (
	"encoding/base64"
	"log"
	"net/http"

	authen "mikel-kunze.com/uploadservice/Authentication"
	logging "mikel-kunze.com/uploadservice/Logging"
	upload "mikel-kunze.com/uploadservice/Upload"
)

func main() {
	correctSetup, err := OnServerStartup()

	if !correctSetup {
		log.Fatal("Failed to start", err)
	}

	// sets the jwt secret. Is needed bevor startup!s
	authen.JWTKey = GetKey()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/auth", sendNewAccess)
	mux.HandleFunc("/api/file/2/", httpFileUploadRequest)

	http.ListenAndServe(":8080", mux)
}

// TODO: check function and refactor
func httpFileUploadRequest(w http.ResponseWriter, r *http.Request) {

	// checks if the client is autorized
	authorized, userName := authen.AuthorizeWithToken(r.Header.Get("Authorization"))

	if !authorized || userName == "" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	logging.LogEntry("[Access]: Upload ", authen.GetIP(r))
	r.ParseMultipartForm(20 << 30)

	if err := upload.HandleUpload(*r.MultipartForm, userName); err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO: check function and refactor
func sendNewAccess(w http.ResponseWriter, r *http.Request) {
	credentials := r.Header.Get("Authorization")
	encoded, err := base64.StdEncoding.DecodeString(credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	isAuthorized, userName := authen.AuthorizeWithOutToken(string(encoded))

	if !isAuthorized {
		w.WriteHeader(http.StatusForbidden)
		return
	} else {
		token, err := authen.GenerateNewAccesstoken(userName)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(token))
		return
	}
}
