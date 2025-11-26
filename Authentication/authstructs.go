package authentication

import "github.com/golang-jwt/jwt/v4"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}
