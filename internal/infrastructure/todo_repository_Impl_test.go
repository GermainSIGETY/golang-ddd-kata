package infrastructure

import (
	"fmt"
	"testing"
	"time"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/stretchr/testify/suite"
)

var (
	toCreateTitle                = "test your repo"
	toCreateDescription          = "test your CRUD"
	toCreateAndReadTitle         = "your repo is funky"
	toCreateAndReadDescription   = "your repo grooves"
	toCreateAndUpdateTitle       = "your repo is fit"
	toCreateAndUpdateDescription = "your repo is strength"
	toCreateAndUpdateAssignee    = "the intern"
)

var toCreateCreationDate = time.Date(2015, 9, 7, 12, 30, 0, 0, time.UTC)
var toCreateDueDate = time.Date(2033, 9, 7, 12, 30, 0, 0, time.UTC)

type updateRequestForTest struct {
	id          int
	title       string
	description string
	dueDate     int64
	assignee    string
}

func (t updateRequestForTest) Id() int {
	return t.id
}

func (t updateRequestForTest) Title() string {
	return t.title
}

func (t updateRequestForTest) Description() string {
	return t.description
}

func (t updateRequestForTest) DueDate() int64 {
	return t.dueDate
}

func (t updateRequestForTest) Assignee() string {
	return t.assignee
}

type InfraTestSuite struct {
	suite.Suite
	repo todosRepository
}

func (suite *InfraTestSuite) SetupSuite() {
	fmt.Println("Setup suite")
	suite.repo = todosRepository{}
	suite.NoError(suite.repo.InitDatabase("/tmp/TodosTest.db", true), "error creating DB")
}

func (suite *InfraTestSuite) TearDownSuite() {
	fmt.Println("Tear down suite")
	suite.repo.db.Migrator().DropTable(&todoGORM{})
	//suite.repo.db.Close()
}

func (suite *InfraTestSuite) TestCreate() {
	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  toCreateDescription,
		DueDate:      toCreateDueDate,
		Title:        toCreateTitle,
		Assignee:     toCreateAndUpdateAssignee,
	}
	id, err := suite.repo.Create(toCreate)
	suite.NoError(err, "error creating Todo")
	suite.NotNil(id, "id is nl")
}

func (suite *InfraTestSuite) TestReadWrongId() {
	todo, err := suite.repo.ReadTodo(676767675776)
	suite.NoError(err)
	suite.Zero(todo)
}

func (suite *InfraTestSuite) TestCreateAndRead() {
	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  toCreateAndReadDescription,
		DueDate:      toCreateDueDate,
		Title:        toCreateAndReadTitle,
	}
	createdId, err := suite.repo.Create(toCreate)
	suite.NoError(err, "error creating Todo")
	suite.NotNil(createdId, "id is nl")

	readTodo, _ := suite.repo.ReadTodo(createdId)

	suite.NotNil(readTodo)
	suite.Equal(createdId, readTodo.ID)
	suite.Equal(toCreateAndReadTitle, readTodo.Title)
	suite.Equal(toCreateAndReadDescription, readTodo.Description)
	suite.Equal(toCreateCreationDate, readTodo.CreationDate)
	suite.Equal(toCreateDueDate, readTodo.DueDate)
}

func (suite *InfraTestSuite) TestCreateAndReadWithEmptyDescription() {
	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  "",
		DueDate:      toCreateDueDate,
		Title:        toCreateAndReadTitle,
	}
	createdId, err := suite.repo.Create(toCreate)
	suite.NoError(err, "error creating Todo")
	suite.NotNil(createdId, "id is nl")

	readTodo, _ := suite.repo.ReadTodo(createdId)

	suite.NotNil(readTodo)
	suite.Equal(createdId, readTodo.ID)
	suite.Empty(readTodo.Description)
}

func (suite *InfraTestSuite) TestCreateAndUpdate() {
	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  toCreateAndUpdateDescription,
		DueDate:      toCreateDueDate,
		Title:        toCreateAndUpdateTitle,
	}
	createdId, errCreate := suite.repo.Create(toCreate)
	suite.NoError(errCreate, "error creating Todo")
	suite.NotNil(createdId, "id is nl")

	readTodo, errRead := suite.repo.ReadTodo(createdId)
	suite.NoError(errRead, "error creating Todo")
	suite.Equal(createdId, readTodo.ID)
	suite.False(readTodo.NotificationSent)

	newTitle := "my test is fat"
	newDescription := "my test is in the ghetto"
	newDueDate := dueDate.Add(time.Hour * 48)
	newDueDateAsInt := newDueDate.Unix()

	updateRequest := updateRequestForTest{createdId, newTitle, newDescription, newDueDateAsInt, assignee}
	api.UpdateFromTodoUpdateRequest(&readTodo, updateRequest)

	readTodo.MarkAsNotificationSent()

	errUpdate := suite.repo.UpdateTodo(readTodo)
	suite.NoError(errUpdate, "error updating todo")

	updatedTodo, errRead2 := suite.repo.ReadTodo(createdId)
	suite.NoError(errRead2, "error reading updated")
	suite.Equal(createdId, updatedTodo.ID)

	suite.Equal(newTitle, updatedTodo.Title)
	suite.Equal(newDescription, updatedTodo.Description)
	suite.Equal(toCreateCreationDate, updatedTodo.CreationDate)
	suite.True(newDueDate.Equal(updatedTodo.DueDate))
	suite.True(updatedTodo.NotificationSent)
}

func (suite *InfraTestSuite) TestReadSummaries() {

	suite.repo.db.Where("1 = 1").Delete(&todoGORM{})

	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  toCreateAndUpdateDescription,
		DueDate:      toCreateDueDate,
		Title:        toCreateAndUpdateTitle,
	}
	id1, err := suite.repo.Create(toCreate)
	suite.NoError(err, "error creating Todo")

	descr2 := "2"
	toCreate2 := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  descr2,
		DueDate:      toCreateDueDate,
		Title:        toCreateAndUpdateTitle,
	}
	_, error2 := suite.repo.Create(toCreate2)
	suite.NoError(error2, "error creating Todo")

	summaries, errRead := suite.repo.ReadTodoList()
	suite.NoError(errRead, "error reading todo list summaries")
	suite.Equal(2, len(summaries))

	mapSummaries := make(map[int]model.ISummaryResponse)
	for _, resp := range summaries {
		mapSummaries[resp.Id()] = resp
	}

	suite.Equal(id1, mapSummaries[id1].Id())
	suite.Equal(toCreateAndUpdateTitle, mapSummaries[id1].Title())
	suite.Equal(toCreateDueDate, mapSummaries[id1].DueDate())
}

func (suite *InfraTestSuite) TestCreateAndDelete() {
	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  toCreateAndUpdateDescription,
		DueDate:      toCreateDueDate,
		Title:        toCreateAndUpdateTitle,
	}
	createdId, err := suite.repo.Create(toCreate)
	suite.NoError(err, "error creating Todo")
	suite.NotNil(createdId, "id is nl")

	deleteError := suite.repo.DeleteTodo(createdId)
	suite.NoError(deleteError, "error deleting Todo")

	deletedRead, readError := suite.repo.ReadTodo(createdId)

	suite.NoError(readError)
	suite.Zero(deletedRead)
}

func (suite *InfraTestSuite) TestDeleteWrongId() {
	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  toCreateAndUpdateDescription,
		DueDate:      toCreateDueDate,
		Title:        toCreateAndUpdateTitle,
	}
	createdId, err := suite.repo.Create(toCreate)
	suite.NoError(err, "error creating Todo")
	suite.NotNil(createdId, "id is nl")

	deleteError := suite.repo.DeleteTodo(99999)
	suite.Error(deleteError, "deleting Todo with incorrect id should return an error")
	suite.Equal(deleteError.Error(), port.NoRowDeleted)

	_, readError := suite.repo.ReadTodo(createdId)
	suite.NoError(readError, "reading deleted error should return error")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestInfraTestSuite(t *testing.T) {
	suite.Run(t, new(InfraTestSuite))
}

func (suite *InfraTestSuite) TestReadTodosIdToNotify() {

	suite.repo.db.Where("1 = 1").Delete(&todoGORM{})

	toCreate := model.Todo{
		ID: 0,

		CreationDate: toCreateCreationDate,
		Description:  toCreateAndUpdateDescription,
		DueDate:      toCreateDueDate,
		Title:        toCreateAndUpdateTitle,
	}
	id1, err := suite.repo.Create(toCreate)
	suite.NoError(err, "error creating Todo")

	descr2 := "2"
	toCreate2 := model.Todo{
		ID: 0,

		CreationDate:     toCreateCreationDate,
		Description:      descr2,
		DueDate:          toCreateDueDate,
		Title:            toCreateAndUpdateTitle,
		NotificationSent: true,
	}
	_, error2 := suite.repo.Create(toCreate2)
	suite.NoError(error2, "error creating Todo")

	ids, errRead := suite.repo.ReadTodosIdsToNotify()
	suite.NoError(errRead, "error reading todo list summaries")
	suite.Equal(1, len(ids))

	suite.Equal(id1, ids[0])
}
