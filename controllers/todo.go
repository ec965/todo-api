package controllers

import (
	"strconv"

	"github.com/ec965/todo-api/models"
	"github.com/gin-gonic/gin"
)

// check if the form has the value
func formHasValue(c *gin.Context, key string) string {
	field := c.PostForm(key)
	if field == "" {
		c.JSON(400, gin.H{"error": key + " is required"})
		return ""
	}
	return field
}

// check if the url has the query param
func queryParamHasValue(c *gin.Context, key string) string {
	field := c.Query(key)
	if field == "" {
		c.JSON(400, gin.H{"error": key + " is required"})
		return ""
	}
	return field
}

// handle database errors
func handleError(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return true
	}
	return false
}

// Create a new todo
func CreateTodo(c *gin.Context) {
	title := formHasValue(c, "title")
	if title == "" {
		return
	}
	todo, err := models.CreateTodo(title)
	if handleError(c, err) {
		return
	}
	c.JSON(200, todo)
}

// Get all todos
func GetTodos(c *gin.Context) {
	todos, err := models.GetAllTodo()
	if handleError(c, err) {
		return
	}
	c.JSON(200, todos)
}

// get a single todo by id query parameter
func GetTodo(c *gin.Context) {
	idStr := queryParamHasValue(c, "id")
	if idStr == "" {
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

func UpdateTodo(c *gin.Context) {
	idStr := formHasValue(c, "id")
	if idStr == "" {
		return
	}
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "id is not a number"})
		return
	}
	id := uint(id64)
	todo, err := models.GetTodoById(id)
	if handleError(c, err) {
		return
	}
	updatedTodo, err := models.UpdateTodo(id, todo.Title, true)
	if handleError(c, err) {
		return
	}
	c.JSON(200, updatedTodo)
}
