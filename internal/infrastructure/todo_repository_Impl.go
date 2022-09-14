package infrastructure

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
)

type todosRepository struct {
	db *gorm.DB
}

const (
	goORMRecordNotFoundError = "record not found"
)

var todoRepository port.ITodosRepository

func InitTodosRepository(URL string, drop bool) error {
	todoRepo := todosRepository{}
	err := todoRepo.InitDatabase(URL, drop)
	todoRepository = todoRepo
	return err
}

func GetTodosRepository() port.ITodosRepository {
	return todoRepository
}

func (repository *todosRepository) InitDatabase(URL string, drop bool) error {
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

func (repository todosRepository) ReadTodoList() ([]model.ISummaryResponse, error) {
	var todos []todoGORM
	err := repository.db.Select("ID, title, due_date").Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return mapToTodoSummaryResponse(todos), nil
}

func mapToTodoSummaryResponse(todos []todoGORM) []model.ISummaryResponse {
	var summaries = make([]model.ISummaryResponse, len(todos))
	for i, t := range todos {
		summaries[i] = model.NewSummaryResponse(t.ID, t.Title, t.DueDate)
	}
	return summaries
}

func (repository todosRepository) ReadTodo(ID int) (model.Todo, error) {
	var todoGORM todoGORM
	err := repository.db.First(&todoGORM, ID).Error
	if err != nil {
		return handleReadError(err)
	}
	return FromTodoGORM(todoGORM), nil
}

func handleReadError(err error) (model.Todo, error) {
	var todo model.Todo
	if err.Error() == goORMRecordNotFoundError {
		return todo, nil
	}
	return todo, err
}

func (repository todosRepository) Create(todo model.Todo) (int, error) {
	var todoGORM = FromTodo(todo)
	db := repository.db.Create(&todoGORM)
	if db.Error != nil {
		return 0, db.Error
	}
	return todoGORM.ID, nil
}

func (repository todosRepository) UpdateTodo(todo model.Todo) error {
	var todoGORM = FromTodo(todo)
	db := repository.db.Save(&todoGORM)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (repository todosRepository) DeleteTodo(id int) error {
	todoGORM := todoGORM{ID: id}
	db := repository.db.Delete(&todoGORM)
	switch {
	case db.Error != nil:
		return db.Error
	case db.RowsAffected == 0:
		return errors.New(port.NoRowDeleted)
	default:
		return nil
	}
}
