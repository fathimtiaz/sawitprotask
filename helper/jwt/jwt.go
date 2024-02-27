package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateUserToken(encodedUser []byte) (string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(""))
	if err != nil {
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["user"] = encodedUser
	claims["exp"] = time.Now()

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}
