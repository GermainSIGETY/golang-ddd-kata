package port

import "github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"

type INotificationSender interface {
	SendToDoNotification(todo model.Todo) error
}
