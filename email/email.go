package email

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/flightlogger/configuration"
)

// Service describes a service capable of sending emails
type Service struct {
	config    configuration.EmailConfiguration
	templates *template.Template
}

// NewEmailService configures a new EmailService
func NewEmailService(config configuration.EmailConfiguration) common.EmailServiceInterface {
	var err error
	s := Service{config: config}

	// Try to read the html-templates
	s.templates, err = template.New("activate").ParseFiles("./email/templates/activateTemplate.html", "./email/templates/resetTemplate.html")

	if err != nil {
		log.Fatalf("Cannot load email templates! %v", err)
	}

	// smtp server configuration.
	return &s
}

func (e *Service) getEmailAuthentication() smtp.Auth {
	log.Println(e.config)
	auth := smtp.PlainAuth("", e.config.Username, e.config.Password, e.config.SmtpServer)
	return auth
}

// sendEmail sets the correct headers and etc
func (e *Service) sendEmail(message string, recipient string, subject string) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectObject := "Subject: " + subject + "!\n"
	msg := []byte(subjectObject + mime + "\n" + message)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", e.config.SmtpServer, e.config.Port), e.getEmailAuthentication(), e.config.Username, []string{recipient}, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// SendResetPasswordLink - Sends an verification email to a client using GMAIL
func (e *Service) SendResetPasswordLink(email string, resetPasswordURL string) error {
	var b bytes.Buffer

	message := fmt.Sprintf(
		`<h3>If you have not requested a new password; Ignore this email</h3>
		<br>
		<h3>To reset your password, use the link below</h3>
		<a href="%s">Your reset link</a>
	`, resetPasswordURL)

	e.templates.ExecuteTemplate(&b, "B", message)

	return e.sendEmail(b.String(), email, "Reset FlightLog password")
}

// SendVerificationEmail - Sends an verification email to a client using GMAIL
func (e *Service) SendVerificationEmail(email string, verificationURL string) error {
	var b bytes.Buffer

	log.Printf("Formatting email for %s, with token %s", email, verificationURL)
	message := fmt.Sprintf(
		`<h3>You want to create a new user on flightlogger. Wise choise</h3>
		<br>
		<h3>To activate your account use the link below</h3>
		<a href="%s">Click this link to activate the better user experience!</a>
	`, verificationURL)

	e.templates.ExecuteTemplate(&b, "B", message)
	return e.sendEmail(b.String(), email, "Activate Flightlog Account")
}
