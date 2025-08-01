package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khanjaved9700/todo_app/middleware"
	"github.com/khanjaved9700/todo_app/models"
	"github.com/khanjaved9700/todo_app/service"
)

type Handler interface {
	CreateTodo(c *gin.Context)
	GetTodoList(c *gin.Context)
	MarkDone(c *gin.Context)
	DeleteTodo(c *gin.Context)
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler(h service.Service) Handler {
	return &handler{service: h}
}

func (h *handler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": "error in binding"})
		return
	}
	var todo = &models.TODO{
		Title: req.Title,
	}

	data, err := h.service.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errpr": "todo not created"})
		return
	}
	c.JSON(http.StatusAccepted, map[string]interface{}{
		"Success": true,
		"Message": "Todo created",
		"Data":    data,
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

func (h *handler) RegisterUser(c *gin.Context) {

	req := &CreateUser{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding failed"})
	}

	hashedPassword, err := middleware.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hashing failed"})
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	userId, err := h.service.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
		"id":      userId,
	})

}

func (s *handler) Login(c *gin.Context) {
	req := &LoginUser{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding failed"})
		return
	}

	user, err := s.service.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Check password
	if !middleware.CheckHashPassword([]byte(user.Password), req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := middleware.GenrateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	// Respond with token
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})

}
