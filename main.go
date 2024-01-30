package main

import (
	"flag"
	"fmt"
	"go-send/graphhelper"
	"log"

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

	// Parse the command-line arguments
	to := flag.String("to", "", "The email address of the recipient")
	from := flag.String("from", "", "The email address of the sender")
	subject := flag.String("subject", "", "The subject of the email")
	message := flag.String("message", "", "The message body of the email")
	flag.Parse()

	// Check if the arguments are valid
	if *to == "" || *from == "" || *subject == "" || *message == "" {
		log.Fatal("Invalid arguments. Please provide To, From, Subject and Message.")
	}

	sendMail(graphHelper, *from, *to, *subject, *message)
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

func sendMail(graphHelper *graphhelper.GraphHelper, sender string, receiver string, subject string, body string) {
	err := graphHelper.SendMail(&subject, &body, sender, &receiver)
	if err != nil {
		log.Panicf("Error sending mail: %v", err)
	}
	fmt.Println("Mail sent.")
	fmt.Println()
}
