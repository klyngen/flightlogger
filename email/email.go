package email

import (
	"fmt"
	"net/smtp"

	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/flightlogger/configuration"
)

// Service describes a service capable of sending emails
type Service struct {
	config configuration.EmailConfiguration
}

// NewEmailService configures a new EmailService
func NewEmailService(config configuration.EmailConfiguration) common.EmailServiceInterface {

	// smtp server configuration.
	fmt.Println("Email Sent!")

	return &Service{config: config}
}

func (e *Service) getEmailAuthentication() smtp.Auth {
	auth := smtp.PlainAuth("FlightLogger", e.config.Username, e.config.Password, e.config.SMTPServer)
	return auth
}

func (e *Service) sendEmail(message string, recipient string) error {
	err := smtp.SendMail(fmt.Sprintf("%s:%s", e.config.SMTPServer, e.config.Port), e.getEmailAuthentication(), e.config.Username, []string{recipient}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// SendResetPasswordLink - Sends an verification email to a client using GMAIL
func (e *Service) SendResetPasswordLink(email string, resetPasswordURL string) error {
	message := fmt.Sprintf(
		`<h3>If you have not requested a new password; Ignore this email</h3>
		<br>
		<h3>To reset your password, use the link below</h3>
		<a href="%s">Your reset link</a>
	`, resetPasswordURL)

	return e.sendEmail(message, email)
}

// SendVerificationEmail - Sends an verification email to a client using GMAIL
func (e *Service) SendVerificationEmail(email string, verificationURL string) error {
	message := fmt.Sprintf(
		`<h3>You want to create a new user on flightlogger. Wise choise</h3>
		<br>
		<h3>To activate your account use the link below</h3>
		<a href="%s">Click this link to activate the better user experience!</a>
	`, verificationURL)

	return e.sendEmail(message, email)
}
