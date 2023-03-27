package port

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
)

type NotificationSenderMock struct {
	Calls int
}

func (mock *NotificationSenderMock) SendToDoNotification(todo model.Todo) error {
	mock.Calls++
	return nil
}

func (mock *NotificationSenderMock) ResetCallsCounter() {
	mock.Calls = 0
}
