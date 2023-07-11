# Kubernetes , Go and React JS with TypeScript
The goal of this project is to build an application with services build using Go in the backend and ReactJS on the frontend.
Why this languages? Well you might have guessed it. They all use types which reduce bugs drastically
This is for learning purposes and therefore critics and comments are allowed :)

## Application architecure.
This application is made up of three major parts.
1. Backend
2. Frontend
3. Kubernetes

### Backend (Golang)
The backend is build using Golang. Golang is my favorite programming language but can be replaced with python.
This is a REST API, and I will use Gin for serving HTTP request.
The API will expose all CRUD operation. These CRUD operations will be performed against Todo items.
A todo items will have the below attributes.
1. Id
2. Title
3. Description
4. State
5. Created at

The backend will be created using the hexagonal architecture. Incase this is not farmiliar to you, kindly research about hexagonal architecture.

### Frontend (React + Typscript)
The frontend is build using ReactJS. Its is my favorite frontend framework at the moment.
The frontend will make HTTP requests to the backend to perform CRUD operation

### Database 
This application will AWS DynamoDB just because its a remote database hence, only database region, access key and secret key are required.
This section can as well be replaced by any other storage.


