package cli

import (
	"flag"
	"go-send/adapters/secondary/graph"
	"go-send/domain/entities"
	"go-send/domain/usecase"
	"os"
)

func RunClient(u *usecase.UserCase) {
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
	m := &entities.Message{
		To:      *to,
		From:    *from,
		Subject: *subject,
		Body:    *message,
	}
	u.SendToGraph(m)
}
