package presentation

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	id    = 1212
	title = "develop tirelessly unit tests"
)

var (
	dueDate = time.Date(1998, time.July, 13, 0, 0, 0, 0, time.UTC)
)

func TestTodo_NewTodoSummaryResponse(t *testing.T) {
	todo := NewTodoSummaryResponse(id, title, dueDate)
	assert.Equal(t, id, todo.Id())
	assert.Equal(t, title, todo.Title())
	assert.Equal(t, dueDate, todo.DueDate())
}
