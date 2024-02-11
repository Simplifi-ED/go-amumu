package smtp

import (
	"errors"
	"fmt"
	"log"

	"github.com/mhale/smtpd"
)

type SMTPSERVER struct{}

func (s *SMTPSERVER) NewSMTPServer(port string, handler smtpd.Handler) {
	host := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Println("SMTP Server is Running...")
	if err := smtpd.ListenAndServe(host, handler, "Happybit", ""); !errors.Is(err, smtpd.ErrServerClosed) {
		log.Fatalf("SMTP server error: %v", err)
	}
	log.Println("Stopped.")

}
