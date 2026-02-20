package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"os"

	_ "embed"

	"gopkg.in/gomail.v2"
)

//go:embed templates/emailVerification.html
var emailTemplateFS embed.FS

func ParseTemplate(filePath string, data any) (string, error) {
	tmpl, err := template.ParseFS(emailTemplateFS, filePath)
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
