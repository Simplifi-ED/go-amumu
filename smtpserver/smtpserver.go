package smtpserver

import (
	"bytes"
	"log"
	"net"
	"net/mail"

	"github.com/mhale/smtpd"
)

func StartSMTPServer() {
	// Start the SMTP server
	log.Println("Starting SMTP server...")
	log.Println("Listening on port 2525...")
	smtpd.ListenAndServe("127.0.0.1:2525", mailHandler, "Happybit", "")
}

func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")
	log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(msg.Body)
	if err != nil {
		return err
	}
	log.Printf("Message: %s", buf.String())
	return nil
}
