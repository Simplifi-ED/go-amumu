package notification

import (
	"go-send/graphhelper"
	. "go-send/utils"
)

type Watcher struct{}

func (w *Watcher) Update(message Message, graphHelper *graphhelper.GraphHelper) {
	SendMail(graphHelper, message.From, message.To, message.Subject, message.Body, false)
}
