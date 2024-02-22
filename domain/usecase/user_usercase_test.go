package usecase

import (
	"errors"
	"testing"

	"go-send/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGraph struct {
	mock.Mock
}

func (m *mockGraph) SendMail(from, to, subject, body string) error {
	args := m.Called(from, to, subject, body)
	return args.Error(0)
}

func TestSendToGraph_Success(t *testing.T) {
	mockGraph := new(mockGraph)
	useCase := NewUserCase(mockGraph)

	message := &entities.Message{
		From:    "from@example.com",
		To:      "to@example.com",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	mockGraph.On("SendMail", message.From, message.To, message.Subject, message.Body).Return(nil)

	err := useCase.SendToGraph(message)

	assert.NoError(t, err)
	mockGraph.AssertExpectations(t)
}

func TestSendToGraph_Error(t *testing.T) {
	mockGraph := new(mockGraph)
	useCase := NewUserCase(mockGraph)

	message := &entities.Message{
		From:    "from@example.com",
		To:      "to@example.com",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	expectedErr := errors.New("error sending mail")
	mockGraph.On("SendMail", message.From, message.To, message.Subject, message.Body).Return(expectedErr)

	err := useCase.SendToGraph(message)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	mockGraph.AssertExpectations(t)
}
