package output

type PortSMTP interface {
	SendEmail(string, string) error
}