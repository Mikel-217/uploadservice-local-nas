package main

import (
	"fmt"
	"log"
	"net/http"

	authen "mikel-kunze.com/uploadservice/Authentication"
	upload "mikel-kunze.com/uploadservice/Upload"
	users "mikel-kunze.com/uploadservice/User"
)

func main() {
	correctSetup, err := OnServerStartup()

	if !correctSetup {
		log.Fatal("Failed to start", err)
	}

	// sets the jwt secret. Is needed bevor startup!
	authen.JWTKey = GetKey()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/2", users.CreateNewUser)
	mux.HandleFunc("/api/users/3", users.DeleteUser)

	mux.HandleFunc("/api/auth", authen.SendNewAccess)

	mux.HandleFunc("/api/file/2", upload.HttpFileUploadRequest)

	fmt.Println("Started successfull")
	fmt.Println(http.ListenAndServe(":8080", mux))
}
