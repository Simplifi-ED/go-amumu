package smtp

import (
	"bytes"
	"errors"
	"fmt"
	"go-send/graphhelper"
	"go-send/utils"
	"log"
	"net"
	"net/mail"

	"github.com/mhale/smtpd"
)

type SMTPSERVER struct {
}

func (s *SMTPSERVER) Init(port string) {
	host := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Println("SMTP Server is Running...")
	if err := smtpd.ListenAndServe(host, mailHandler, "Happybit", ""); !errors.Is(err, smtpd.ErrServerClosed) {
		log.Fatalf("SMTP server error: %v", err)
	}
	log.Println("Stopped.")

}

func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
	graphHelper := graphhelper.NewGraphHelper()

	utils.InitializeGraph(graphHelper)

	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")
	log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
	buf := new(bytes.Buffer)
	buf.ReadFrom(msg.Body)
	log.Printf("Message: %s", buf.String())
	utils.SendMail(graphHelper, from, to[0], subject, buf.String(), false)
	return nil
}
