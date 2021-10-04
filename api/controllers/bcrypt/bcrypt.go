package bcrypt

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type PasswordBcrypt struct {
	Password string `json:"password" validate:"required,min=8"`
}

func CryptPassword(w http.ResponseWriter, r *http.Request) {
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
