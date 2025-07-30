package service

import (
	"github.com/khanjaved9700/todo_app/models"
	"github.com/khanjaved9700/todo_app/repository"
)

type Service interface {
	CreateTodo(title string) error
	GetTodoList() ([]models.TODO, error)
	MarkDone(id uint) error
	DeleteTodo(id uint) error
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateTodo(title string) error {
	req := &models.TODO{
		Title: title,
	}
	if err := s.repository.CreateTodo(req); err != nil {
		return err
	}
	return nil

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
