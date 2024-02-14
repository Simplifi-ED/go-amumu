package usecase

import (
	"flag"
	"go-send/adapters/primary/server"
	"go-send/infrastructure/notification"
	"os"

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
	smtpserver := server.NewSmtpServer()
	config := server.NewConfig(*serverPort)
	backend := &server.Backend{
		Channel: ch,
	}
	if err := smtpserver.RunSMTPServer(config, backend); err != nil {
		log.Fatal(err)
	}

}
