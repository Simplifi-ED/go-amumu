package main

import (
	"flag"
	"fmt"
	"go-send/graphhelper"
	. "go-send/notification"
	"go-send/smtp"
	. "go-send/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var subjectmq *Subject

func main() {
	fmt.Println("Go SMTP <=> Graph")

	loadEnv()

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

	flag.Usage = func() {
		fmt.Println("\nUsage:")
		fmt.Printf(" %s [subcommand] [options]\n", os.Args[0])
		fmt.Println("server options:")
		fmt.Printf(" -port string\n")
		fmt.Println("client options:")
		fmt.Printf(" -to string - The email address of the recipient\n")
		fmt.Printf(" -from string - The email address of the sender\n")
		fmt.Printf(" -subject string - The subject of the email\n")
		fmt.Printf(" -message string - The message body of the email\n")
		fmt.Printf(" -channel boolean - Send to MS Teams channel (default=false)\n")
	}

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println("Initializing App...")
	graphHelper := graphhelper.NewGraphHelper()
	InitializeGraph(graphHelper)

	subjectmq = &Subject{}
	subjectmq.SetGraphHelper(graphHelper)
	watcher := &Watcher{}

	smtpHandler := smtp.NewMailHandler(subjectmq, graphHelper)
	subjectmq.Register(watcher)

	switch os.Args[1] {
	case "server":
		ServerCmd.Parse(os.Args[2:])
		smtpserver := smtp.SMTPSERVER{}
		smtpserver.NewSMTPServer(*serverPort, smtpHandler.HandleEmail)
	case "client":
		ClientCmd.Parse(os.Args[2:])
		// Check if the arguments are valid
		if *to == "" || *from == "" || *subject == "" || *message == "" {
			flag.Usage()
		}

		SendMail(graphHelper, *from, *to, *subject, *message, *channel)
	default:
		flag.Usage()
		os.Exit(1)
	}

}

func loadEnv() {
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}
