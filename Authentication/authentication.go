package authentication

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	database "mikel-kunze.com/uploadservice/Database"
)

// TODO: get this also out of conf.json
var jwtKey = []byte("")

func AuthorizeWithToken(token string) (bool, string) {

	strings.Replace(token, "Baerer", "", 0)

	// to validate JWT
	claims := Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil || !tkn.Valid {
		return false, ""
	}

	// searches the token in the DB
	if !database.GetToken(token) {
		return false, ""
	}

	tokenClaims := tkn.Claims.(*Claims)

	if tokenClaims.Username == "" {
		return false, ""
	}

	return true, tokenClaims.Username
}

// If the User has no token
func AuthorizeWithOutToken(authData string) (bool, string) {

	credentials := strings.Split(authData, ";")
	user := database.GetUserByName(credentials[0])

	if user.UserName == credentials[0] && user.PW == credentials[1] {
		return true, user.UserName
	}

	return false, ""
}

// Generates a new JWT for the user to login faster --> gets send to the frontend --> sets the cookie there
func GenerateNewAccesstoken(username string) (string, error) {
	expiraionTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiraionTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(jwtKey)

	database.SetNewToken(tokenString, expiraionTime)

	return tokenString, err
}
