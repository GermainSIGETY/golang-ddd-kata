package infrastructure

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"os"
)

type mailSender struct {
	mailgun *mailgun.MailgunImpl
	domain  string
}

func NewNotificationSender() port.INotificationSender {
	sender := mailSender{}
	apikey := os.Getenv("MAILGUN_PRIVATE_API_KEY")
	sender.domain = os.Getenv("MAILGUN_DOMAIN")
	sender.mailgun = mailgun.NewMailgun(sender.domain, apikey)
	return sender
}

func (m mailSender) SendToDoNotification(todo model.Todo) error {
	message := m.composeMessage(todo)
	resp, id, err := m.mailgun.Send(context.Background(), message)
	if err != nil {
		log.Err(err).Msg("")
		return err
	}
	log.Info().Str("ID", id).Str("response", resp).Msg("")
	return nil
}

func (m mailSender) composeMessage(todo model.Todo) *mailgun.Message {

	sender := "mailgun@" + m.domain
	subject := todo.Title
	body := "What you have to do is : " + todo.Description
	recipient := todo.Assignee

	// The message object allows you to add attachments and Bcc recipients
	return m.mailgun.NewMessage(sender, subject, body, recipient)
}
