package database

import "time"

// TODO: implement Database functions
// - func for storring accesstokens
// - func for storring userdata

func SetNewToken(token string, expiration time.Time) {}

func GetUserByName(userName string) UserStruct { return UserStruct{} }

func GetToken(token string) bool { return false }
