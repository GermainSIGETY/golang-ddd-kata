package todo

import "time"

type SummaryResponse struct {
	id      int
	title   string
	dueDate time.Time
}

func NewSummaryResponse(id int, title string, dueDate time.Time) SummaryResponse {
	return SummaryResponse{id, title, dueDate}
}

func (t SummaryResponse) Id() int {
	return t.id
}

func (t SummaryResponse) Title() string {
	return t.title
}

func (t SummaryResponse) DueDate() time.Time {
	return t.dueDate
}
