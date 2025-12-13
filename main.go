package main

import (
	"fmt"
	"log"
	"net/http"

	authen "mikel-kunze.com/uploadservice/Authentication"
	directorys "mikel-kunze.com/uploadservice/FileHandling/Directory"
	files "mikel-kunze.com/uploadservice/FileHandling/Files"

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

	// User Requests
	mux.HandleFunc("/api/users/2", users.CreateNewUser) // creates a new User
	mux.HandleFunc("/api/users/3", users.DeleteUser)    // deletes a User

	// Authentication requests
	mux.HandleFunc("/api/auth", authen.SendNewAccess)

	// File requests
	mux.HandleFunc("/api/file/2", files.HttpFileUploadRequest) // for uploading
	mux.HandleFunc("/api/file/3", files.HttpFileUploadRequest) // deletes a files

	// Directory requests
	mux.HandleFunc("/api/dir/2", directorys.HttpDirRequest) // for creating dirs
	mux.HandleFunc("/api/dir/3", directorys.HttpDirRequest) // for deleting dirs

	fmt.Println("Started successfull")
	fmt.Println(http.ListenAndServe(":8080", mux))
}
