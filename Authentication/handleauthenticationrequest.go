package authentication

import (
	"encoding/base64"
	"net/http"
)

// Sends a new authentication token to the user
func SendNewAccess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	credentials := r.Header.Get("Authorization")
	encoded, err := base64.StdEncoding.DecodeString(credentials)

	if err != nil || string(encoded) == "" {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	isAuthorized, userName := AuthorizeWithOutToken(string(encoded))

	if !isAuthorized {
		w.WriteHeader(http.StatusForbidden)
		return
	} else {
		token, err := GenerateNewAccesstoken(userName)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(token))
		return
	}
}
