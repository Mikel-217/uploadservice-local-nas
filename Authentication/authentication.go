package authentication

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	database "mikel-kunze.com/uploadservice/database"
	"mikel-kunze.com/uploadservice/startup"
)

// The secret
var JWTKey = startup.GetKey()

// Checks for a valide jwt token and if the token is saved in the database
// Returns the bool and the claims
func AuthorizeWithToken(token string) (bool, *Claims) {

	strings.Replace(token, "Baerer", "", 0)

	// to validate JWT
	claims := Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { return JWTKey, nil })

	if err != nil || !tkn.Valid {
		return false, &Claims{}
	}

	// searches the token in the DB
	if !database.CheckTokenExistence(token) {
		return false, &Claims{}
	}

	tokenClaims := tkn.Claims.(*Claims)

	if tokenClaims.Username == "" {
		return false, &Claims{}
	}

	return true, tokenClaims
}

// If the User has no token --> he gets a new one
func AuthorizeWithOutToken(authData string) (bool, database.UserStruct) {

	credentials := strings.Split(authData, ";")
	user := database.GetUserByName(credentials[0])

	if user.UserName == credentials[0] && user.PW == credentials[1] {
		return true, user
	}

	return false, database.UserStruct{}
}

// Generates a new JWT for the user to login faster --> gets send to the frontend --> sets the cookie there
func GenerateNewAccesstoken(user database.UserStruct) (string, error) {
	expiraionTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:        user.ID,
		Username:      user.UserName,
		UserDirectory: "DEFAULT_DIR",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiraionTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(JWTKey)

	// sets new accesstoken into db
	database.CreateNewToken(tokenString, expiraionTime)

	return tokenString, err
}
