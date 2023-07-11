package http

import (
	"fmt"

	"github.com/AntonyIS/k8-go-react-typscript-todo/backend/config"
	"github.com/AntonyIS/k8-go-react-typscript-todo/backend/core"
	"github.com/AntonyIS/k8-go-react-typscript-todo/backend/internal/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitHTTPRoutes(svc *core.HTTPHandlerService, c config.Config) {
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
	todoroutes := router.Group("/api/v1/todo")

	handler := core.NewHTTPService(svc)

	{
		todoroutes.GET("/", handler.GetUsers)
		todoroutes.GET("/:id", handler.GetUser)
		todoroutes.POST("/", handler.PostUser)
		todoroutes.PUT("/:id", handler.PutUser)
		todoroutes.DELETE("/:id", handler.DeleteUser)
	}
	port := fmt.Sprintf(":%s", c.Port)

	router.Run(port)
}
