package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
)

// SMTP configuration including recipient address
var smtpConfig = map[string]string{
	"host":      "smtp.gmail.com",
	"port":      "587",
	"user":      "user@yourpool.org",
	"password":  "password",
	"toAddress": "info@yourpool.org",
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Daten aus dem Request lesen
	var requestData map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Error decoding request data", http.StatusBadRequest)
		return
	}

	// Set the recipient address from SMTP configuration
	toAddress := smtpConfig["toAddress"]
	fromAddress := smtpConfig["user"] // use SMTP user as sender address
	subject := requestData["subject"]
	body := requestData["body"]

	// Hier die E-Mail senden
	err = sendEmail(toAddress, fromAddress, subject, body)
	if err != nil {
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	// Erfolgreiche Antwort an den Client senden
	response := map[string]string{"message": fmt.Sprintf("E-Mail erfolgreich gesendet an %s mit Betreff '%s'", toAddress, subject)}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func sendEmail(toAddress, fromAddress, subject, body string) error {
	auth := smtp.PlainAuth("", smtpConfig["user"], smtpConfig["password"], smtpConfig["host"])
	smtpAddr := smtpConfig["host"] + ":" + smtpConfig["port"]

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\r\n", fromAddress, toAddress, subject, body)

	return smtp.SendMail(smtpAddr, auth, fromAddress, []string{toAddress}, []byte(msg))
}

func main() {
	http.HandleFunc("/send-email", sendEmailHandler)

	fmt.Println("Server gestartet auf http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
