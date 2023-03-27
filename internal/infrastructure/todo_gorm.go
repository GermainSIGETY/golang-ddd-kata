package infrastructure

import (
	"time"
)

// why a pointer only on description field ?
// Because only this field is not mandatory and can be null in database
// -in Datatabase we want to register column with null value if Description is no set
// -in Code we use we Golang's zero value of a string which is empty string if Description is not set
type todoGORM struct {
	ID               int       `gorm:"primary_key"`
	Title            string    `gorm:"Column:title;size:255"`
	Description      *string   `gorm:"Column:description;size:255"`
	CreationDate     time.Time `gorm:"Column:creation_date"`
	DueDate          time.Time `gorm:"Column:due_date"`
	Assignee         *string   `gorm:"Column:assignee;size:255"`
	NotificationSent bool      `gorm:"Column:notificationSent"`
}

func (todoGORM) TableName() string {
	return "Todo"
}
