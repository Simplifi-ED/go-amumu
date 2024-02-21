package graph

import (
	"os"

	"github.com/charmbracelet/log"
)

type IGraphEmail interface {
	SendMail(sender string, receiver string, subject string, body string)
}

type GraphEmail struct {
	Graph *GraphHelper
}

func (g *GraphEmail) SendMail(sender string, receiver string, subject string, body string) {
	log.Info("Sending Mail...")
	err := g.Graph.Send(&subject, &body, &sender, &receiver)
	if err != nil {
		log.Error("Error sending message", "Error:", err)
		os.Exit(1)
	}
	log.Info("Mail sent.")
}
