package server

import (
	"go-send/domain/entities"
	"go-send/infrastructure/notification"
	"strings"
	"testing"

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

func TestSession_Data(t *testing.T) {
	// Initialize testify assert
	assert := assert.New(t)

	// Prepare test data
	testMessage := `From: sender@example.com
To: receiver@example.com
Subject: Test Subject
This is a test email body.`

	mockChannel := &notification.Subject{}
	mockChannel.SetGraphEmail(&mockGraph{})
	session := &Session{channel: mockChannel}

	// Prepare a mock reader with the test data
	r := strings.NewReader(testMessage)

	// Call the function being tested
	err := session.Data(r)

	// Assert that no error occurred during the execution
	assert.NoError(err, "Data() function returned an unexpected error")

	// Check if the message has been properly set in the channel
	expectedMessage := entities.Message{
		From:    "sender@example.com",
		To:      "receiver@example.com",
		Subject: "Test Subject",
		Body:    "This is a test email body.",
	}

	assert.Equal(expectedMessage, mockChannel.Message, "Data() did not set the message correctly in the channel")
}
