package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khanjaved9700/todo_app/handlers"
	"github.com/khanjaved9700/todo_app/repository"
	"github.com/khanjaved9700/todo_app/service"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewHandler(services)

	api := r.Group("/todo")

	api.POST("/create", handler.CreateTodo)
	api.GET("/list", handler.GetTodoList)
	api.PUT("/done/:id", handler.MarkDone)
	api.DELETE("/delete/:id", handler.DeleteTodo)
}
