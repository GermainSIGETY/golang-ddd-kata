package infrastructure

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func SkipTestSendEmail(t *testing.T) {
	todo := model.Todo{
		ID:           12,
		CreationDate: time.Date(2015, 9, 7, 12, 30, 0, 0, time.UTC),
		Description:  "this is a task that has to be done !",
		DueDate:      time.Date(2025, 9, 7, 12, 30, 0, 0, time.UTC),
		Title:        "a task",
		Assignee:     "bob@unknownEmailDomainThatDoNotExists.fr",
	}
	mailSender := NewNotificationSender()
	mailError := mailSender.SendToDoNotification(todo)
	assert.NoError(t, mailError)

}
