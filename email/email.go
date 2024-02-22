package email

import (
	"log"
	"net/smtp"
	"strings"
)

func SendAlert(title string, description string, character string) {
	auth := smtp.PlainAuth("", EMAIL_ADDRESS, EMAIL_APP_PASSWORD, EMAIL_AUTH_HOST)

	to := strings.Split(strings.Trim(EMAIL_RECIPIENTS_CSL, " "), ",")
	msg := []byte("Subject: New personailty just dropped\r\n" +
		"\r\n" +
		"Gamers! A new Ryan Gosling movie just dropped and you know what that means\r\n" +
		"There's a new personailty to download!\r\n" +
		"\r\n" +
		"Our main man is playing " + character + " in '" + title + "', described as:\r\n" +
		"'" + description + "'\r\n" +
		"\r\n" +
		"Enjoy your new personality, friend :)\r\n")

	err := smtp.SendMail(EMAIL_SEND_MAIL_ADDRESS, auth, EMAIL_ADDRESS, to, msg)

	if err != nil {
		log.Fatal(err)
	}
}
