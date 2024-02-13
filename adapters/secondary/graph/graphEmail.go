package graph

import (
	"github.com/charmbracelet/log"
)

type GraphEmail struct {
	Graph *GraphHelper
}

func (g *GraphEmail) SendMail(sender string, receiver string, subject string, body string) {
	log.Info("Sending Mail...")
	err := g.Graph.Send(&subject, &body, &sender, &receiver)
	if err != nil {
		log.Error("Error sending message: %v", err)
	}
	log.Info("Mail sent.")
}
