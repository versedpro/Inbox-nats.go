package mailing

import (
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

// Sender sends mails.
type Sender interface {
	Send(to, subject, body string) error
}

type sender struct {
	from mail.Address
	addr string
	auth smtp.Auth
}

// NewSender implementation using an SMTP server.
func NewSender(from, host, port, username, password string) Sender {
	return &sender{
		from: mail.Address{Address: from},
		addr: net.JoinHostPort(host, port),
		auth: smtp.PlainAuth("", username, password, host),
	}
}

func (s *sender) Send(to, subject, body string) error {
	toAddr := mail.Address{Address: to}

	var msg string
	msg += fmt.Sprintf("From: %s\r\n", s.from.String())
	msg += fmt.Sprintf("To: %s\r\n", toAddr.String())
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += "Content-Type: text/html; charset=utf-8\r\n"
	msg += "\r\n"
	msg += body

	return smtp.SendMail(
		s.addr,
		s.auth,
		s.from.Address,
		[]string{toAddr.Address},
		[]byte(msg),
	)
}
