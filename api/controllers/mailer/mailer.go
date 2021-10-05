package mailer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
)

func SendMail(w http.ResponseWriter, r *http.Request) {

	// Sender data.
	from := "......@gmail.com"
	password := "********"

	// Receiver email address.
	to := []string{
		"........@.......com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Mail Sended !",
	})
}
