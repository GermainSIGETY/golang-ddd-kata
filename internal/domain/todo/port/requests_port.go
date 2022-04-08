package port

type UpdateRequest interface {
	Id() int
	Title() string
	Description() string
	DueDate() int64
}

type CreationRequest interface {
	Title() string
	Description() string
	DueDate() int64
}
