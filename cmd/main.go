package main

import (
	"flag"
	"fmt"
	"go-send/graphhelper"
	"go-send/smtp"
	"go-send/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go SMTP<=>Graph")
	fmt.Println()

	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	//server commands set
	ServerCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := ServerCmd.String("port", "2525", "port")

	//client commands set
	ClientCmd := flag.NewFlagSet("client", flag.ExitOnError)
	to := ClientCmd.String("to", "", "The email address of the recipient")
	from := ClientCmd.String("from", "", "The email address of the sender")
	subject := ClientCmd.String("subject", "", "The subject of the email")
	message := ClientCmd.String("message", "", "The message body of the email")
	channel := ClientCmd.Bool("channel", false, "Send to MS Teams channel")

	if len(os.Args) < 2 {
		fmt.Println("expected 'server' or 'client' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		ServerCmd.Parse(os.Args[2:])
		smtpserver := smtp.SMTPSERVER{}
		smtpserver.Init(*serverPort)
	case "client":
		ClientCmd.Parse(os.Args[2:])
		// Check if the arguments are valid
		if *to == "" || *from == "" || *subject == "" || *message == "" {
			fmt.Println("Invalid arguments. Please provide the following arguments:")
			log.Fatal("-to | -from | -subject | -message")
		}
		graphHelper := graphhelper.NewGraphHelper()
		utils.InitializeGraph(graphHelper)
		utils.SendMail(graphHelper, *from, *to, *subject, *message, *channel)
	default:
		fmt.Println("expected 'server' or 'client' subcommands")
		os.Exit(1)
	}

}
