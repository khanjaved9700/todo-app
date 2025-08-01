package service

import (
	"github.com/khanjaved9700/todo_app/models"
	"github.com/khanjaved9700/todo_app/repository"
)

type Service interface {
	CreateTodo(req *models.TODO) (*models.TODO, error)
	GetTodoList() ([]models.TODO, error)
	MarkDone(id uint) error
	DeleteTodo(id uint) error
	RegisterUser(req *models.User) (uint, error)
	GetUserByEmail(email string) (models.User, error)
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateTodo(req *models.TODO) (*models.TODO, error) {

	data, err := s.repository.CreateTodo(req)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (s *service) GetTodoList() ([]models.TODO, error) {
	return s.repository.GetTodoList()
}

func (s *service) MarkDone(id uint) error {
	return s.repository.MarkDone(id)
}

func (s *service) DeleteTodo(id uint) error {
	return s.repository.Delete(id)
}

func (s *service) RegisterUser(req *models.User) (uint, error) {

	userId, err := s.repository.RegisterUser(req)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *service) GetUserByEmail(email string) (models.User, error) {

	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
