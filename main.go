package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/rs/cors"
)

// EmailData struct holds the structure of email data
type EmailData struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/send-email", sendEmailHandler)

	// Use the Cors module for CORS support
	handler := cors.Default().Handler(http.DefaultServeMux)

	port := 3000
	fmt.Printf("Server is running on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var emailData EmailData
	err := json.NewDecoder(r.Body).Decode(&emailData)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// You need to adjust the SMTP settings here
	smtpHost := "your_smtp_server"
	smtpPort := 587
	smtpUsername := "your_username"
	smtpPassword := "your_password"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Prepare email content
	message := []byte(
		fmt.Sprintf("To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n", emailData.To, emailData.Subject, emailData.Message),
	)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpHost, smtpPort),
		auth,
		smtpUsername,
		[]string{emailData.To},
		message,
	)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	fmt.Println("Email sent successfully.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}
