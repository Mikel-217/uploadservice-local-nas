package database

import "time"

type UserStruct struct {
	ID       uint   // Is also included in the JWT
	UserName string // Is also included in the JWT
	PW       string // !important: always hash this string
}

type UserDirectorys struct {
	DirID   uint
	UserID  uint // foreign key for User
	DirName string
	DirPath string
}

type ActiveAccessTokens struct {
	TokenID        uint
	ActiveToken    string
	ExpirationDate time.Time
}
