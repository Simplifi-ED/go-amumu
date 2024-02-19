package server

import (
	"context"
	"log"

	"github.com/emersion/go-smtp"
)

type SmtpServer struct {
	host *smtp.Server
}

func NewSmtpServer(config *Config, backend *Backend) *SmtpServer {
	server := smtp.NewServer(backend)
	server.Addr = config.Address
	server.Domain = config.Domain
	server.WriteTimeout = config.WriteTimeout
	server.ReadTimeout = config.ReadTimeout
	server.MaxMessageBytes = config.MaxMessageBytes
	server.MaxRecipients = config.MaxRecipients
	server.AllowInsecureAuth = config.AllowInsecureAuth
	return &SmtpServer{
		host: server,
	}
}

func (srv SmtpServer) RunSMTPServer() error {

	log.Println("Starting SMTP server at", srv.host.Addr)
	return srv.host.ListenAndServe()
}

func (srv SmtpServer) ShutdownSMTPServer(ctx context.Context) error {
	return srv.host.Shutdown(ctx)
}
