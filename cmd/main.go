package main

import (
	"flag"
	"fmt"
	"go-send/graphhelper"
	"go-send/smtp"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Graph App-Only")
	fmt.Println()

	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	graphHelper := graphhelper.NewGraphHelper()

	initializeGraph(graphHelper)

	//server commands set
	ServerCmd := flag.NewFlagSet("server", flag.ExitOnError)
	// serverPort := ServerCmd.String("port", "2525", "port")

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
		smtpserver.Init("2525")
	case "client":
		ClientCmd.Parse(os.Args[2:])
		// Check if the arguments are valid
		if *to == "" || *from == "" || *subject == "" || *message == "" {
			fmt.Println("Invalid arguments. Please provide the following arguments:")
			log.Fatal("-to | -from | -subject | -message")
		}

		sendMail(graphHelper, *from, *to, *subject, *message, *channel)
	default:
		fmt.Println("expected 'server' or 'client' subcommands")
		os.Exit(1)
	}

}

func initializeGraph(graphHelper *graphhelper.GraphHelper) {
	err := graphHelper.InitializeGraphForAppAuth()
	if err != nil {
		log.Panicf("Error initializing Graph for app auth: %v\n", err)
	}
}

func displayAccessToken(graphHelper *graphhelper.GraphHelper) {
	token, err := graphHelper.GetAppToken()
	if err != nil {
		log.Panicf("Error getting user token: %v\n", err)
	}

	fmt.Printf("App-only token: %s", *token)
	fmt.Println()
}

func listUsers(graphHelper *graphhelper.GraphHelper) {
	users, err := graphHelper.GetUsers()
	if err != nil {
		log.Panicf("Error getting users: %v", err)
	}

	// Output each user's details
	for _, user := range users.GetValue() {
		fmt.Printf("User: %s\n", *user.GetDisplayName())
		fmt.Printf("  ID: %s\n", *user.GetId())

		noEmail := "NO EMAIL"
		email := user.GetMail()
		if email == nil {
			email = &noEmail
		}
		fmt.Printf("  Email: %s\n", *email)
	}

	// If GetOdataNextLink does not return nil,
	// there are more users available on the server
	nextLink := users.GetOdataNextLink()

	fmt.Println()
	fmt.Printf("More users available? %t\n", nextLink != nil)
	fmt.Println()
}

func sendMail(graphHelper *graphhelper.GraphHelper, sender string, receiver string, subject string, body string, channel bool) {
	err := graphHelper.SendMail(&subject, &body, sender, &receiver, channel)
	if err != nil {
		log.Panicf("Error sending message: %v", err)
	}
	fmt.Println("Mail sent.")
	fmt.Println()
}
