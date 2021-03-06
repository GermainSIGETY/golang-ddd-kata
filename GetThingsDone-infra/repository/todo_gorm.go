package repository

import (
	"time"
)

type TodoGORM struct {
	ID           *int      `gorm:"primary_key"`
	Title        string    `gorm:"Column:title;size:255"`
	Description  *string   `gorm:"Column:description;size:255"`
	CreationDate time.Time `gorm:"Column:creation_date"`
	DueDate      time.Time `gorm:"Column:due_date"`
}

func (TodoGORM) TableName() string {
	return "Todo"
}
