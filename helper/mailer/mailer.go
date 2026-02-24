package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

//go:embed templates/emailVerification.html
var emailVerificationFS embed.FS

//go:embed templates/passwordReset.html
var passwordResetFS embed.FS

func ParseTemplate(fs embed.FS, filePath string, data any) (string, error) {
	tmpl, err := template.ParseFS(fs, filePath)
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
		emailVerificationFS,
		"templates/emailVerification.html",
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

func SendPasswordReset(target, passwordResetLink string) error {
	message := gomail.NewMessage()

	body, err := ParseTemplate(
		passwordResetFS,
		"templates/passwordReset.html",
		map[string]any{
			"passwordResetLink": passwordResetLink,
		},
	)

	if err != nil {
		return err
	}

	message.SetHeader("From", os.Getenv("MAILER_USER"))
	message.SetHeader("To", target)
	message.SetHeader("Subject", "Password Reset")
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		"smtp.gmail.com",         // SMTP host
		587,                      // SMTP port
		os.Getenv("MAILER_USER"), // username
		os.Getenv("MAILER_PASS"), // password
	)

	return dialer.DialAndSend(message)
}
