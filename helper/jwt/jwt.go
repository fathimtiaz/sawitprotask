package jwt

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateUserToken(userId int64, userPhone, userName string) (string, error) {
	// TODO store in config
	privateKeyBytes, err := ioutil.ReadFile(os.Getenv("PRIV_KEY"))
	if err != nil {
		log.Fatal("Error reading private key:", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyBytes))
	if err != nil {
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["user_id"] = userId
	claims["user_phone"] = userPhone
	claims["user_name"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
}

func ParseToken(token string) (claims jwt.MapClaims, err error) {
	var parsed *jwt.Token
	var ok bool

	// TODO store in config
	publicKeyBytes, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		log.Fatal("Error reading public key:", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyBytes))
	if err != nil {
		return nil, err
	}

	parsed, err = jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return publicKey, nil
	})

	if !parsed.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok = parsed.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("invalid token")
	}

	return
}
