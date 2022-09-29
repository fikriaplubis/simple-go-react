package todos

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	todo, status, err := service.CreateTodos(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, todo)

}

func TestGetTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	service.CreateTodos(req)

	todos, status, err := service.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, false, todos[0].Done)
	assert.Equal(t, req.Task, todos[0].Task)

	req = DataRequest{
		Task: "task 2",
	}

	service.CreateTodos(req)

	todos, status, err = service.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 2, len(todos))
	assert.Equal(t, false, todos[1].Done)
	assert.Equal(t, req.Task, todos[1].Task)

}

func TestUpdateTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	newReq := DataRequest{
		Task: "task 1",
	}

	newTodo, _, _ := service.CreateTodos(newReq)

	updatedReq := DataRequest{
		Task: "task 2",
	}

	updatedTodo, status, err := service.UpdateTodos(int(newTodo.ID), updatedReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, updatedTodo)
}

func TestDeleteService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	newReq := DataRequest{
		Task: "task 1",
	}

	todo, _, _ := service.CreateTodos(newReq)

	status, err := service.DeleteTodos(int(todo.ID))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
}
