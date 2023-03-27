package services

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/rs/zerolog/log"
	"net/mail"
	"time"
)

type notificationService struct {
	repository          port.ITodosRepository
	notificationSender  port.INotificationSender
	notificationChannel chan int
	tickerTime          int
}

// NewNotificationService Create NotificationService and only return notificationChannel:
// to send a notification, send a Todo Id to the channel and notificationService will do the job
func NewNotificationService(repository port.ITodosRepository, notificationSender port.INotificationSender, tickerTime int) chan<- int {
	service := notificationService{
		repository:          repository,
		notificationSender:  notificationSender,
		notificationChannel: make(chan int, 100),
		tickerTime:          tickerTime,
	}
	go service.receiveFromChannel()
	service.failureRecovery()
	return service.notificationChannel
}

func (notificationService *notificationService) receiveFromChannel() {
	for {
		select {
		case todoId := <-notificationService.notificationChannel:
			notificationService.loadTodoAndNotify(todoId)
		}
	}
}

func (notificationService *notificationService) loadTodoAndNotify(todoId int) {
	log.Info().Int("todoId", todoId).Msg("notification to Send for Todo")
	if todo, err := notificationService.repository.ReadTodo(todoId); err != nil {
		log.Err(err).Msg("cannot read todo, notification not sent")
	} else if !todo.NotificationSent {
		notificationService.sendNotification(todo)
	}
}

func (notificationService *notificationService) sendNotification(todo model.Todo) {
	if todo.NotificationSent {
		log.Info().Int("todoId", todo.ID).Msg(" notification already sent, skip")
		return
	}
	if err := notificationService.sendNotificationIfAssigneeIsValid(todo); err != nil {
		log.Err(err).Msg("notification not sent")
	} else {
		notificationService.persistTodoWithNotificationSent(todo)
	}
}

func (notificationService *notificationService) sendNotificationIfAssigneeIsValid(todo model.Todo) error {
	if emailValid(todo.Assignee) {
		return notificationService.notificationSender.SendToDoNotification(todo)
	} else {
		log.Info().Str("email", todo.Assignee).Msg("notification skipped because email is invalid")
		return nil
	}
}

func emailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (notificationService *notificationService) persistTodoWithNotificationSent(todo model.Todo) {
	todo.MarkAsNotificationSent()
	if err := notificationService.repository.UpdateTodo(todo); err != nil {
		log.Err(err).Msg("cannot persistTodoWithNotificationSent")
	}
}

func (notificationService *notificationService) failureRecovery() {
	ticker := time.NewTicker(time.Duration(notificationService.tickerTime) * time.Second)
	go func() {
		for {
			select {
			case _ = <-ticker.C:
				log.Info().Msg("try to notify Todos")
				notificationService.notifyUnotifiedTodos()
			}
		}
	}()
}

func (notificationService *notificationService) notifyUnotifiedTodos() {
	if ids, err := notificationService.repository.ReadTodosIdsToNotify(); err != nil {
		log.Err(err).Msg("cannot ReadTodosIdsToNotify")
	} else {
		for _, id := range ids {
			notificationService.notificationChannel <- id
		}
	}
}
