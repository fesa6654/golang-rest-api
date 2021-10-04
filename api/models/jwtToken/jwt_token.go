package jwt_token_model

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWTKey() []byte {
	var jwtKey = []byte("secret_key")
	return jwtKey
}

func NewJWTToken(username string) string {

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTKey())

	if err != nil {
		fmt.Println("token can not created !")
	}

	return tokenString
}
