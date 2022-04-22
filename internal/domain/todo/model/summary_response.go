package model

import "time"

// ISummaryResponse Immutability pattern : object cannot be modified after its creation by constructor
type ISummaryResponse interface {
	Id() int
	Title() string
	DueDate() time.Time
}

type summaryResponse struct {
	id      int
	title   string
	dueDate time.Time
}

func NewSummaryResponse(id int, title string, dueDate time.Time) ISummaryResponse {
	// L'intérêt de spécifié les key+value lors de l'instanciation d'une struct, c'est de
	// - ne pas être dépendant de l'ordre des champs de la struct
	// - rendre le code plus clair et lisible
	return summaryResponse{
		id:      id,
		title:   title,
		dueDate: dueDate,
	}
}

func (t summaryResponse) Id() int {
	return t.id
}

func (t summaryResponse) Title() string {
	return t.title
}

func (t summaryResponse) DueDate() time.Time {
	return t.dueDate
}
