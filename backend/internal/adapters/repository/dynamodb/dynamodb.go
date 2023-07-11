package dynamodb

import (
	"errors"
	"fmt"

	"github.com/AntonyIS/k8-go-react-typscript-todo/backend/config"
	"github.com/AntonyIS/k8-go-react-typscript-todo/backend/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type TodoDynamoDBClient struct {
	db        *dynamodb.DynamoDB
	tablename string
}

func NewTodoDynamoDBClient(c config.Config) *TodoDynamoDBClient {
	creds := credentials.NewStaticCredentials(c.AWSAccessKey, c.AWSAccessSecretKey, "")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(c.AWSRegion),
		Credentials: creds,
	}))
	return &TodoDynamoDBClient{
		db:        dynamodb.New(sess),
		tablename: c.TableName,
	}
}

func (t *TodoDynamoDBClient) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	// Convert todo to a dynamodb attributes
	entityParsed, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return nil, err
	}

	// Prepare input to be added into the database
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(t.tablename),
	}
	// Add todo into the database
	_, err = t.db.client.PutItem(input)

	if err != nil {
		return nil, err
	}

	return todo, nil

}

func (t *TodoDynamoDBClient) ReadTodo(id string) (*domain.Todo, error) {
	// Read todo from the database with id
	result, err := t.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(t.tablename),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}
	// Check if todo exists
	if result.Item == nil {
		return nil, errors.New(fmt.Sprintf("todo with id %s not found", id))
	}
	// Initialize todo type
	var todo *domain.Todo
	// Convert todo to todo type from dynamodb type
	err = dynamodbattribute.UnmarshalMap(result.Item, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil

}

func (t *TodoDynamoDBClient) ReadTodos() (*[]domain.Todo, error) {
	// Create the input parameters for the Scan operation
	params := &dynamodb.ScanInput{
		TableName: aws.String(t.tablename), // Replace with your table name
	}
	// Perform the Scan operation
	result, err := t.db.Scan(params)
	if err != nil {
		return nil, err
	}
	var todos *[]domain.Todo
	// Process the scan results
	for _, item := range result.Items {
		// create todo instance
		var todo = *domain.Todo{}
		todo.id = item["Id"]
		todo.title = item["title"]
		todo.description = item["description"]
		todo.State = item["State"]
		todo.createdAt = item["createdAt"]

		// add todo to todos list
		todo = append(todos, todo)
	}
	return &todos, nil
}

func (t *TodoDynamoDBClient) UpdateTodo(todo *domain.Todo) (*domain.Todo, error) {
	// marshal todo to dynamodb attributes
	entityParsed, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(t.tablename),
	}

	_, err = t.db.PutItem(input)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *TodoDynamoDBClient) DeleteTodo(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(t.db.tablename),
	}

	res, err := t.db.DeleteItem(input)
	if res != nil {
		return err
	}
	return nil
}
