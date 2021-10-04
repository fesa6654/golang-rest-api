package jwt_token

import (
	"encoding/json"
	"net/http"
	"strings"

	token "golang-rest-api/api/models/jwtToken"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(w http.ResponseWriter, r *http.Request) {

	tokenString := token.NewJWTToken("Sefa ÜN")

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

}

func CheckJWTToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Authorization") == "" {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Token not found",
		})
		return
	}

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if reqToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Token",
		})
		return
	}

	claims := &token.Claims{}

	tkn, err := jwt.ParseWithClaims(reqToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return token.JWTKey(), nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid Token",
			})
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Token",
		})
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Token",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"kullanıcı": claims.Username,
		"status":    "Token Kontrol Başarılı",
	})
}
