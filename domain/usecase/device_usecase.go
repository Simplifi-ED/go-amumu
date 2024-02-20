package usecase

import (
	"context"
	"flag"
	"go-send/adapters/primary"
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
	var configPath string
	ServerCmd.StringVar(&configPath, "config", "/etc/amumu-config.yaml", "path to config file")

	if err := primary.ValidateConfigPath(configPath); err != nil {
		log.Fatal("Error validating file: %v", "Error:", err)
	}
	ServerCmd.Parse(os.Args[2:])
	config, err := server.NewConfig(configPath)
	if err != nil {
		log.Fatal("Error loading config", "Error", err)
	}
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
