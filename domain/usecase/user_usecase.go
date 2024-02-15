package usecase

import (
	"go-send/adapters/secondary/graph"
	"go-send/domain/entities"
)

type UserCase struct {
	gh *graph.GraphEmail
}

func NewUserCase(graph *graph.GraphEmail) *UserCase {
	return &UserCase{
		gh: graph,
	}
}

func (usecase *UserCase) SendToGraph(message *entities.Message) error {
	usecase.gh.SendMail(message.From, message.To, message.Subject, message.Body)
	return nil
}
