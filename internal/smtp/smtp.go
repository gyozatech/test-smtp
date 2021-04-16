package smtp

// Sender interface represents a email sender
// object that can send a message
type Sender interface {
	Send(from string, to []string, msg string) error
}

// NewSender construct a new sender for a given hostname
func NewSender(hostname string) Sender {
	return &sender{
		Hostname: hostname,
	}
}
