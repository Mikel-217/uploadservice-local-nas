package database

import "time"

type UserStruct struct {
	ID       uint   // Is also inclided in the JWT
	UserName string // Is also included in the JWT
	PW       string // !important: always hash this string
	Dirs     UserDirectorys
}

type UserDirectorys struct {
	DirName string
	DirPath string
}

type ActiveAccessTokens struct {
	ActiveToken    string
	ExpirationDate time.Time
}
