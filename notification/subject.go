package notification

import (
	"go-send/graphhelper"
)

type Subject struct {
	observers   []Observer
	message     Message
	graphhelper *graphhelper.GraphHelper
}

func (s *Subject) SetGraphHelper(gh *graphhelper.GraphHelper) {
	s.graphhelper = gh
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
		observer.Update(s.message, s.graphhelper)
	}
}

func (s *Subject) SetEvent(message Message) {
	s.message = message
	s.Notify()
}
