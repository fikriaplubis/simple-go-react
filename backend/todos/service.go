package todos

import (
	"net/http"

	"github.com/fikriaplubis/simple-go-react/models"
)

type Service interface {
	GetTodos() ([]models.Todos, int, error)
	CreateTodos(req DataRequest) (models.Todos, int, error)
	UpdateTodos(id int, req DataRequest) (models.Todos, int, error)
	DeleteTodos(id int) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetTodos() ([]models.Todos, int, error) {

	todos, err := s.repo.GetTodos()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return todos, http.StatusOK, nil
}

func (s *service) CreateTodos(req DataRequest) (models.Todos, int, error) {

	todo, err := s.repo.CreateTodos(req.Task)
	if err != nil {
		return models.Todos{}, http.StatusInternalServerError, err
	}

	return todo, http.StatusOK, nil
}

func (s *service) UpdateTodos(id int, req DataRequest) (models.Todos, int, error) {

	todo, err := s.repo.UpdateTodos(id, req)
	if err != nil {
		return models.Todos{}, http.StatusInternalServerError, err
	}

	return todo, http.StatusOK, nil
}

func (s *service) DeleteTodos(id int) (int, error) {
	err := s.repo.DeleteTodos(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
