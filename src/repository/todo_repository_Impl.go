package repository

import (
	"errors"
	"github.com/GermainSIGETY/golang-ddd-kata/src/domain"
	"github.com/GermainSIGETY/golang-ddd-kata/src/infrastructure"
	"github.com/GermainSIGETY/golang-ddd-kata/src/presentation"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodosRepository struct {
	db *gorm.DB
}

const (
	goORMRecordNotFoundError = "record not found"
)

func (repository *TodosRepository) InitDatabase(URL string, drop bool) error {
	var err error

	dialect2 := sqlite.Open(URL)

	repository.db, err = gorm.Open(dialect2)
	//repository.db.LogMode(true)
	if err != nil {
		return err
	}
	if drop {
		migrator := repository.db.Migrator()
		migrator.DropTable(&todoGORM{})
	}
	repository.db.AutoMigrate(&todoGORM{})
	return nil
}

func (repository TodosRepository) ReadTodoList() ([]presentation.TodoSummaryResponse, error) {
	var todos []todoGORM
	err := repository.db.Select("ID, title, due_date").Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return mapToTodoSummaryResponse(todos), nil
}

func mapToTodoSummaryResponse(todos []todoGORM) []presentation.TodoSummaryResponse {
	var summaries = make([]presentation.TodoSummaryResponse, len(todos))
	for i, todo := range todos {
		summaries[i] = presentation.NewTodoSummaryResponse(*todo.ID, todo.Title, todo.DueDate)
	}
	return summaries
}

func (repository TodosRepository) ReadTodo(ID int) (domain.Todo, error) {
	var todoGORM todoGORM
	err := repository.db.First(&todoGORM, ID).Error
	if err != nil {
		return handleReadError(err)
	}
	return FromTodoGORM(todoGORM), nil
}

func handleReadError(err error) (domain.Todo, error) {
	var todo domain.Todo
	if err.Error() == goORMRecordNotFoundError {
		return todo, nil
	}
	return todo, err
}

func (repository TodosRepository) Create(todo domain.Todo) (int, error) {
	var todoGORM = FromTodo(todo)
	db := repository.db.Create(&todoGORM)
	if db.Error != nil {
		return 0, db.Error
	}
	return *todoGORM.ID, nil
}

func (repository TodosRepository) UpdateTodo(todo domain.Todo) error {
	var todoGORM = FromTodo(todo)
	db := repository.db.Save(&todoGORM)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (repository TodosRepository) DeleteTodo(ID int) error {
	todoGORM := todoGORM{ID: &ID}
	db := repository.db.Delete(&todoGORM)
	switch {
	case db.Error != nil:
		return db.Error
	case db.RowsAffected == 0:
		return errors.New(infrastructure.NoRowDeleted)
	default:
		return nil
	}
}
