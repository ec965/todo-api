package routers

import (
	"github.com/ec965/todo-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", controllers.Ping)

	todo := router.Group("/todo")
	todo.POST("", controllers.CreateTodo)
	todo.GET("/all", controllers.GetTodos)
	todo.GET("", controllers.GetTodo)

	return router
}
