package main

import (
	"flag"
	"fmt"
	"log"

	"go-send/graphhelper"
	"go-send/smtpserver"

	"github.com/joho/godotenv"

)

func main() {
	fmt.Println("Go Graph App-Only")
	fmt.Println()

	// Load .env files
	loadEnv()

	// Parse the command-line arguments
	to, from, subject, message, channel, server := parseFlags()

	// initializeGraph(graphHelper)
	// Check if the -server flag is provided
	if *server {
		smtpserver.StartSMTPServer()
		// Check if the arguments are valid
	} else if *to == "" || *from == "" || *subject == "" || *message == "" {
		log.Fatal("Invalid arguments. Please provide To, From, Subject and Message.")
		// if no server flag provided and the other flag are valid send email
	} else if !*server {
		graphHelper := graphhelper.NewGraphHelper()
		initializeGraph(graphHelper)
		sendMail(graphHelper, *from, *to, *subject, *message, *channel)
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

func loadEnv() {
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func parseFlags() (*string, *string, *string, *string, *bool, *bool) {
	to := flag.String("to", "", "The email address of the recipient")
	from := flag.String("from", "", "The email address of the sender")
	subject := flag.String("subject", "", "The subject of the email")
	message := flag.String("message", "", "The message body of the email")
	channel := flag.Bool("channel", false, "Send to MS Teams channel")
	server := flag.Bool("server", false, "Start SMTP server")
	flag.Parse()
	return to, from, subject, message, channel, server
}