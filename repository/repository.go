package repository

import (
	"github.com/khanjaved9700/todo_app/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTodo(todo *models.TODO) (*models.TODO, error)
	GetTodoList() ([]models.TODO, error)
	MarkDone(id uint) error
	Delete(id uint) error
	RegisterUser(req *models.User) (uint, error)
	GetUserByEmail(email string) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateTodo(todo *models.TODO) (*models.TODO, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
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

func (r *repository) RegisterUser(req *models.User) (uint, error) {

	if err := r.db.Create(&req).Error; err != nil {
		return 0, err
	}

	return req.ID, nil
}

func (r *repository) GetUserByEmail(email string) (models.User, error) {

	var user models.User

	if err := r.db.Where("email=?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
