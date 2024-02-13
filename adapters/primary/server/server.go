package server

import (
	"log"

	"github.com/emersion/go-smtp"
)

type SmtpServer struct{}

func NewSmtpServer() *SmtpServer {
	return &SmtpServer{}
}

func (srv SmtpServer) RunSMTPServer(config *Config, backend *Backend) error {
	server := smtp.NewServer(backend)
	server.Addr = config.Address
	server.Domain = config.Domain
	server.WriteTimeout = config.WriteTimeout
	server.ReadTimeout = config.ReadTimeout
	server.MaxMessageBytes = config.MaxMessageBytes
	server.MaxRecipients = config.MaxRecipients
	server.AllowInsecureAuth = config.AllowInsecureAuth

	log.Println("Starting SMTP server at", server.Addr)
	return server.ListenAndServe()
}
