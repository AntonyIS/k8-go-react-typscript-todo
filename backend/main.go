package main

import (
	"example.com/todo-be/config"
	"example.com/todo-be/internal/adapters/app"
	"example.com/todo-be/internal/adapters/repository/dynamodb"
	"example.com/todo-be/internal/core"
)

func main() {
	// application service
	config := config.Config{
		AWSRegion:          "me-south-1",
		AWSAccessKey:       "top secret",
		AWSAccessSecretKey: "top top super secret",
		Port:               "5000",
	}
	// Application Dynamodb client
	rep := dynamodb.NewTodoDynamoDBClient(config)
	// Application/Todo service
	svc := core.NewTodoService(&rep)
	// Application HTTP service
	app.InitHTTPRoutes(*svc)

}
