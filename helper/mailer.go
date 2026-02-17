package helper

import (
	"bytes"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

func ParseTemplate(templatePath string, data any) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}

func SendEmailVerification(target, emailVerificationLink string) error {
	message := gomail.NewMessage()

	body, err := ParseTemplate(
		"view/emailVerification.html",
		map[string]any{
			"emailVerificationLink": emailVerificationLink,
		},
	)

	if err != nil {
		return err
	}

	message.SetHeader("From", os.Getenv("MAILER_USER"))
	message.SetHeader("To", target)
	message.SetHeader("Subject", "Email Verification")
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		"smtp.gmail.com",         // SMTP host
		587,                      // SMTP port
		os.Getenv("MAILER_USER"), // username
		os.Getenv("MAILER_PASS"), // password
	)

	return dialer.DialAndSend(message)
}
