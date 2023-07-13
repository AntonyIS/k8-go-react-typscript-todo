package main

import (
	"fmt"

	"example.com/todo-be/config"
	"example.com/todo-be/internal/adapters/app"
	"example.com/todo-be/internal/adapters/repository/dynamodb"
	"example.com/todo-be/internal/core"
)

func init() {
	config.LoadEnv()
}
func main() {
	// application configuration
	config := config.NewConfiguration()
	fmt.Println(config)
	// Application Dynamodb client
	rep := dynamodb.NewTodoDynamoDBClient(config)
	// Application/Todo service
	svc := core.NewTodoService(&rep)
	// Application HTTP Routes
	app.InitHTTPRoutes(*svc, config.Port)

}
