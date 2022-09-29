package api

import (
	"github.com/fikriaplubis/simple-go-react/todos"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	todosRepo := todos.NewRepository(s.DB)
	todosService := todos.NewService(todosRepo)
	todosHandler := todos.NewHandler(todosService)

	s.Router.GET("/", todosHandler.GetTodos)
	s.Router.POST("/send", todosHandler.CreateTodo)
	s.Router.PUT("/:id", todosHandler.UpdateTodos)
	s.Router.DELETE("/:id", todosHandler.DeleteTodos)
}
