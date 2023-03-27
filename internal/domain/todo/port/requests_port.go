package port

type UpdateRequest interface {
	Id() int
	Title() string
	Description() string
	DueDate() int64
	Assignee() string
}

type CreationRequest interface {
	Title() string
	Description() string
	DueDate() int64
	Assignee() string
}
