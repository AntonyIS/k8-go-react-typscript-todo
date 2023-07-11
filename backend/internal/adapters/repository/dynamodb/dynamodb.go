package dynamodb

import (
	"errors"
	"fmt"

	"example.com/todo-be/config"
	"example.com/todo-be/internal/core"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type TodoDynamoDBClient struct {
	db        *dynamodb.DynamoDB
	tablename string
}

func NewTodoDynamoDBClient(c config.Config) core.TodoRepository {
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

func (t *TodoDynamoDBClient) CreateTodo(todo *core.Todo) (*core.Todo, error) {
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
	_, err = t.db.PutItem(input)

	if err != nil {
		return nil, err
	}

	return todo, nil

}

func (t *TodoDynamoDBClient) ReadTodo(id string) (*core.Todo, error) {
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
	var todo *core.Todo
	// Convert todo to todo type from dynamodb type
	err = dynamodbattribute.UnmarshalMap(result.Item, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil

}

func (t *TodoDynamoDBClient) ReadTodos() (*[]core.Todo, error) {
	todos := []core.Todo{}
	filt := expression.Name("Id").AttributeNotExists()
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("title"),
		expression.Name("description"),
		expression.Name("state"),
		expression.Name("createdAt"),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if err != nil {
		return nil, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(t.tablename),
	}
	result, err := t.db.Scan(params)

	if err != nil {
		return nil, err
	}

	for _, item := range result.Items {
		var todo core.Todo

		err = dynamodbattribute.UnmarshalMap(item, &todo)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)

	}
	return &todos, nil
}

func (t *TodoDynamoDBClient) UpdateTodo(todo *core.Todo) (*core.Todo, error) {
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
		TableName: aws.String(t.tablename),
	}

	res, err := t.db.DeleteItem(input)
	if res != nil {
		return err
	}
	return nil
}
