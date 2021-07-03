package controllers

import (
	"github.com/ec965/todo-api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Create a new todo
func CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.JSON(400, gin.H{"error": "title is required"})
		return
	}
	todo, err := models.CreateTodo(title)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, todo)
}

// Get all todos
func GetTodos(c *gin.Context) {
	todos, err := models.GetAllTodo()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, todos)
}

// get a single todo by id query parameter
func GetTodo(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "id is not a number"})
		return
	}
	id := uint(id64)
	todo, err := models.GetTodoById(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todo)
}
