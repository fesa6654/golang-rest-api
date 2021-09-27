package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username" validate:"required,email,min=5,max=100"`
	Password string `json:"password" validate:"required,min=8"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Username and Password not found !",
		})
		return
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Password Error !",
		})
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error !",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token":  tokenString,
		"status": "OK",
	})
}

func TokenControl(w http.ResponseWriter, r *http.Request) {
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

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(reqToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
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

type PasswordBcrypt struct {
	Password string `json:"password" validate:"required,min=8"`
}

func BcryptPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var passwordBcrypted PasswordBcrypt
	err := json.NewDecoder(r.Body).Decode(&passwordBcrypted)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Password not found !",
		})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordBcrypted.Password), 8)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Password Bcrypt Error !",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"encrypedPassword": string(bytes),
		"message":          "Password Bcrypted !",
	})
}

type PasswordDecrypt struct {
	Hash     string `json:"hash" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=8"`
}

func DecryptPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var passwordDecrypted PasswordDecrypt
	err := json.NewDecoder(r.Body).Decode(&passwordDecrypted)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Password and Hash not found !",
		})
		return
	}

	passwordStatus := bcrypt.CompareHashAndPassword([]byte(passwordDecrypted.Hash), []byte(passwordDecrypted.Password))

	if passwordStatus != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Decrypt Error !",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Decrypt is Running !",
	})
}
