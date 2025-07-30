package repository

import (
	"github.com/khanjaved9700/todo_app/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTodo(todo *models.TODO) error
	GetTodoList() ([]models.TODO, error)
	MarkDone(id uint) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateTodo(todo *models.TODO) error {

	if err := r.db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTodoList() ([]models.TODO, error) {
	var todo []models.TODO
	if err := r.db.Find(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *repository) MarkDone(id uint) error {

	var todo models.TODO

	if err := r.db.First(&todo, id).Error; err != nil {
		return err
	}
	if err := r.db.Model(&todo).Update("done", true).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&models.TODO{}, id).Error; err != nil {
		return err
	}
	return nil
}
