package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/juanpicasti/go-todo-app/app/customerror"
	"github.com/juanpicasti/go-todo-app/app/dtos"
	"github.com/juanpicasti/go-todo-app/app/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Create(c *gin.Context) {
	var requestBody dtos.TodoCreateRequest

	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, ok := c.Get("UserID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get user ID from context."})
		return
	}

	newTodo, err := h.service.Create(requestBody, userId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func (h *TodoHandler) Update(c *gin.Context) {
	var requestBody dtos.TodoCreateRequest

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo id must be provided as a positive integer."})
		return
	}

	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := h.service.Update(requestBody, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Todo with the given ID not found: %d", id)})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

func (h *TodoHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo ID must be provided as a positive integer."})
		return
	}

	todoResponse, err := h.service.GetById(id)

	var notFounderror *customerror.TodoNotFoundError

	if errors.As(err, &notFounderror) {
		c.JSON(http.StatusNotFound, gin.H{"error": notFounderror.Error()})
		return
	} else if err != nil {
		log.Error().Err(err).Msg(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something unexpected happened when trying to get the todo."})
		return
	}

	c.JSON(http.StatusOK, todoResponse)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo ID must be provided as a positive integer."})
		return
	}

	todoResponse, err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo with the given ID not found."})
		return
	}

	c.JSON(http.StatusOK, todoResponse)
}
