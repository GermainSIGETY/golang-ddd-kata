package model

import (
	"time"
)

// Todo The entity, which is by nature mutable; domain performs operations on it, compute stuff, change values etc, and in some case store it after operations
type Todo struct {
	// ID is an important field. So it must be at the top of the struct fields list
	ID int

	// Others fields must be sort alphabetically to easily find a field, when we read the code
	CreationDate     time.Time
	Description      string
	DueDate          time.Time
	Title            string
	Assignee         string
	NotificationSent bool
}

// MapToTodoResponse is currently overkill because it is strictly identical than Todo Entity
// but if an entity has:
//   - kind of internal field
//   - fields in another format than ui (ui has more ready to serve formats)
//
// we should need theses 'responses' struct in order to have a 'struct contract' used by ui, independent from domain (aka shock absorber pattern)
func (t *Todo) MapToTodoResponse() ReadTodoResponse {
	if t.ID == 0 {
		return ReadTodoResponse{}
	}
	return ReadTodoResponse{
		ID:           t.ID,
		Title:        t.Title,
		Description:  t.Description,
		CreationDate: t.CreationDate,
		DueDate:      t.DueDate,
		Assignee:     t.Assignee,
	}
}

func (t *Todo) MarkAsNotificationSent() {
	t.NotificationSent = true
}
