package notification

import (
	"go-send/adapters/secondary/graph"
	"go-send/domain/entities"
)

type Watcher struct{}

func (w *Watcher) HandleEvent(message entities.Message, graphEmail *graph.GraphEmail) {
	graphEmail.SendMail(message.From, message.To, message.Subject, message.Body)
}
