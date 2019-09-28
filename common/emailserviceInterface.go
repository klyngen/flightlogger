package common

// EmailServiceInterface describes a service capable of sending verification-emails
type EmailServiceInterface interface {
	SendResetPasswordLink(email string, resetPasswordURL string) error
	SendVerificationEmail(email string, verificationURL string) error
}
