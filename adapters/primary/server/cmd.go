package server

import (
	"flag"
	"log"
	"os"
)

func RunServer() {
	ServerCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := ServerCmd.String("port", "2525", "port")
	ServerCmd.Parse(os.Args[2:])

	config := NewConfig(*serverPort)
	backend := &Backend{}
	smtpserver := NewSmtpServer()
	if err := smtpserver.RunSMTPServer(config, backend); err != nil {
		log.Fatal(err)
	}
}
