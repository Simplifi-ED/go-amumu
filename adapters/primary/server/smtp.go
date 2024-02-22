package server

import (
	"go-send/domain/entities"
	"go-send/infrastructure/notification"
	"io"
	"net/mail"
	"strings"

	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type Backend struct {
	Channel *notification.Subject
}

// NewSession is called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{
		channel: bkd.Channel,
	}, nil
}

// A Session is returned after successful login.
type Session struct {
	channel *notification.Subject
}

// AuthPlain implements authentication using SASL PLAIN.
func (s *Session) AuthPlain(username, password string) error {
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {

	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {

	return nil
}

func (s *Session) Data(r io.Reader) error {
	if msg, err := io.ReadAll(r); err != nil {
		return err
	} else {
		msg, err := mail.ReadMessage(strings.NewReader(string(msg)))
		if err != nil {
			return err
		}

		from := msg.Header.Get("From")
		to := msg.Header.Get("To")
		subject := msg.Header.Get("Subject")
		body, err := io.ReadAll(msg.Body)
		if err != nil {
			return err
		}
		m := &entities.Message{
			To:      to,
			From:    from,
			Subject: subject,
			Body:    string(body),
		}
		s.channel.SetEvent(*m)
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
