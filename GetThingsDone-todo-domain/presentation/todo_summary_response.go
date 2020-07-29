package presentation

import "time"

type TodoSummaryResponse struct {
	id      int
	title   string
	dueDate time.Time
}

func NewTodoSummaryResponse(id int, title string, dueDate time.Time) TodoSummaryResponse {
	return TodoSummaryResponse{id, title, dueDate}
}

func (t TodoSummaryResponse) Id() int {
	return t.id
}

func (t TodoSummaryResponse) Title() string {
	return t.title
}

func (t TodoSummaryResponse) DueDate() time.Time {
	return t.dueDate
}
