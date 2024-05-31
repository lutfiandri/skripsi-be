package helper

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"skripsi-be/internal/config"
)

func SendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + config.SmtpName + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	log.Println(config.SmtpEmail)
	log.Println(config.SmtpPassword)
	log.Println(config.SmtpName)
	log.Println(config.SmtpHost)
	log.Println(config.SmtpPort)

	auth := smtp.PlainAuth("", config.SmtpEmail, config.SmtpPassword, config.SmtpHost)
	smtpAddr := fmt.Sprintf("%s:%s", config.SmtpHost, config.SmtpPort)

	err := smtp.SendMail(smtpAddr, auth, config.SmtpEmail, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
