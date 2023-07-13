package app

import (
	"fmt"

	"example.com/todo-be/internal/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitHTTPRoutes(svc core.TodoApplicationService, port string) {
	// Enable detailed error responses
	gin.SetMode(gin.DebugMode)

	// Setup Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Group users API
	todoroutes := router.Group("/api/v1/todos")

	handler := core.NewHTTPService(svc)

	{
		todoroutes.GET("/", handler.GetTodos)
		todoroutes.GET("/:id", handler.GetTodo)
		todoroutes.POST("/", handler.PostTodo)
		todoroutes.PUT("/:id", handler.UpdateTodo)
		todoroutes.DELETE("/:id", handler.DeleteTodo)
	}
	port = fmt.Sprintf(":%s", port)

	router.Run(port)
}
