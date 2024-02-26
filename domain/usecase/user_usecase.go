package usecase

import (
	"go-send/adapters/secondary/graph"
	"go-send/domain/entities"

	"github.com/charmbracelet/log"
)

type IUserCase interface {
	SendToGraph(message *entities.Message) error
}

type UserCase struct {
	gh graph.IGraphEmail
}

func NewUserCase(graph graph.IGraphEmail) IUserCase {
	return &UserCase{
		gh: graph,
	}
}

func (usecase *UserCase) SendToGraph(message *entities.Message) error {
	log.Info("Sending Mail...")
	err := usecase.gh.SendMail(message.From, message.To, message.Subject, message.Body)
	if err != nil {
		log.Error("Error sending message", "Error:", err)
		return err
	}
	log.Info("Mail sent.")
	return nil
}
