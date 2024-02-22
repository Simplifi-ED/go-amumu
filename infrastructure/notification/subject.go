package notification

import (
	"go-send/adapters/secondary/graph"
	"go-send/domain/entities"
)

type Subject struct {
	observers  []Observer
	Message    entities.Message
	graphEmail graph.IGraphEmail
}

func (s *Subject) SetGraphEmail(gh graph.IGraphEmail) {
	s.graphEmail = gh
}

func (s *Subject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Deregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.HandleEvent(s.Message, s.graphEmail)
	}
}

func (s *Subject) SetEvent(message entities.Message) {
	s.Message = message
	s.Notify()
}
