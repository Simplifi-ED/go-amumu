package smtp

import (
	"bytes"
	"go-send/graphhelper"
	"go-send/notification"
	"log"
	"net"
	"net/mail"
)

type MailHandler struct {
	Channel *notification.Subject
	helper  *graphhelper.GraphHelper
}

func NewMailHandler(ch *notification.Subject, graphHelper *graphhelper.GraphHelper) *MailHandler {
	return &MailHandler{
		Channel: ch,
		helper:  graphHelper,
	}
}

func (m *MailHandler) HandleEmail(origin net.Addr, from string, to []string, data []byte) error {
	msg, err := mail.ReadMessage(bytes.NewReader(data))
	if err != nil {
		log.Printf("failed to read message: %v", err)
	}
	subject := msg.Header.Get("Subject")
	log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
	buf := new(bytes.Buffer)
	buf.ReadFrom(msg.Body)
	log.Printf("Message: %s", buf.String())
	message := &notification.Message{
		To:      to[0],
		From:    from,
		Body:    buf.String(),
		Subject: subject,
	}
	m.Channel.SetEvent(*message)
	return nil
}
