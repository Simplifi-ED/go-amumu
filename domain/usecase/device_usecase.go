package usecase

import (
	"context"
	"flag"
	"go-send/adapters/primary/server"
	"go-send/infrastructure/notification"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
)

type DeviceCase struct{}

func NewDeviceCase() *DeviceCase {
	return &DeviceCase{}
}

func (d *DeviceCase) RunServer(ch *notification.Subject) {
	ServerCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := ServerCmd.String("port", "2525", "port")
	ServerCmd.Parse(os.Args[2:])
	config := server.NewConfig(*serverPort)
	backend := &server.Backend{
		Channel: ch,
	}
	smtpserver := server.NewSmtpServer(config, backend)

	go func() {
		if err := smtpserver.RunSMTPServer(); err != nil {
			log.Fatal("HTTP server error: %v", "Error:", err)
		}
		log.Info("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := smtpserver.ShutdownSMTPServer(shutdownCtx); err != nil {
		log.Fatal("SMTP shutdown error: %v", "Error:", err)
	}
	log.Info("Graceful shutdown complete.")

}
