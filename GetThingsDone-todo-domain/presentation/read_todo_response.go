package presentation

import "time"

type ReadTodoResponse struct {
	ID           int
	Title        string
	Description  *string
	CreationDate time.Time
	DueDate      time.Time
}
