package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khanjaved9700/todo_app/service"
)

type Handler interface {
	CreateTodo(c *gin.Context)
	GetTodoList(c *gin.Context)
	MarkDone(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler(h service.Service) Handler {
	return &handler{service: h}
}

func (h *handler) CreateTodo(c *gin.Context) {
	var todo struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": "error in binding"})
		return
	}

	if err := h.service.CreateTodo(todo.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": "todo not created"})
		return
	}
	c.JSON(http.StatusAccepted, map[string]interface{}{
		"success": "ok",
	})

}

func (h *handler) GetTodoList(c *gin.Context) {
	todoList, err := h.service.GetTodoList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": "data not fetched"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "ok",
		"todo":    todoList,
	})
}
func (h *handler) MarkDone(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := h.service.MarkDone(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": "list not marked done"})
		return
	}
	c.JSON(http.StatusAccepted, map[string]interface{}{
		"success": "ok",
	})

}
func (h *handler) DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := h.service.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": "list not marked done"})
		return
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"mresult": "successfully deleted",
	})
}
