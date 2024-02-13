package cli

import (
	"flag"
	"go-send/adapters/secondary/graph"
	"os"
)

func RunClient(gh *graph.GraphHelper) {
	ClientCmd := flag.NewFlagSet("client", flag.ExitOnError)
	to := ClientCmd.String("to", "", "The email address of the recipient")
	from := ClientCmd.String("from", "", "The email address of the sender")
	subject := ClientCmd.String("subject", "", "The subject of the email")
	message := ClientCmd.String("message", "", "The message body of the email")
	channel := ClientCmd.Bool("channel", false, "Send to MS Teams channel")
	ClientCmd.Parse(os.Args[2:])

	if *to == "" || *from == "" || *subject == "" || *message == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *channel == true {
		graphTeams := &graph.GraphTeams{}
		graphTeams.SendAlert(*message)
	}
	gh.Send(subject, message, from, to)
}
