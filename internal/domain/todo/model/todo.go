package model

import (
	"time"
)

// Soit transformer cette struct en immutable avec une interface définsissant TOUTES les methodes dispo de l'exterieur sur cette struct ( y compris les "Getters")
// Soit mettre les champs en externe pour éviter les getters inutiles ( Je préfère cette solution )
type Todo struct {
	// ID is an important field so it must be a the top of the struct fields list
	ID int

	// Others fields must be sort alphabetically to easly find a field, when we read the code
	CreationDate time.Time
	Description  string
	DueDate      time.Time
	Title        string
}

// MapToTodoResponse actuellement "inutile" puisque simplement passe plat, mais si l'app grandi, cette fonction et nécessaire pour que le domain soit en accord avec son "contrat" de struct de donnée avec l'ui qui l'utilise
func (t Todo) MapToTodoResponse() ReadTodoResponse {
	if t.ID == 0 {
		return ReadTodoResponse{}
	}
	return ReadTodoResponse{
		ID:           t.ID,
		Title:        t.Title,
		Description:  t.Description,
		CreationDate: t.CreationDate,
		DueDate:      t.DueDate,
	}
}
