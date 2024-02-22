package notification

import (
	"go-send/adapters/secondary/graph"
	"go-send/domain/entities"
)

type Observer interface {
	HandleEvent(entities.Message, graph.IGraphEmail)
}
