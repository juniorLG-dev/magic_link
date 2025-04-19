package input

type PortService interface {
	SendEmail(string) error
	VerifyCode(string) error
}