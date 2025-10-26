# Taskmanager
## Task Manager API(Golang + GIN + JWT)
A complete RESTful API built with **Go (Gin framework)** for managing tasks (Create, Read, Update, Delete) with optional **JWT authentication** and **Ngrok deployment** for public access.

## FOLDER STRUCTURE
  task-manager/
│
├── main.go
├── go.mod
├── go.sum
│
├── models/
│   ├── task.go
│   └── user.go
│
├── repository/
│   ├── task_repository.go
│   └── user_repository.go
│
├── handlers/
│   ├── task_handler.go
│   └── auth_handler.go
│
├── routes/
│   └── routes.go
│
├── middleware/
│   └── auth.go
│
└── utils/
    └── jwt.go

---

## Tech Stack
 Language: Go (Golang)
 Framework: Gin Gonic
 Auth: JWT (JSON Web Tokens)
 Deployment: Ngrok

## Getting Started

##  Clone the Repository

git clone https://github.com/RITIK0012/task-manager.git
cd task-manager


## Install Dependencies
go mod tidy

## Run the Server
go run main.go


## Available Routes

| Method     | Endpoint     | Description               |
| ---------- | ------------ | ------------------------- |
| **POST**   | `/tasks`     | Create a new task         |
| **GET**    | `/tasks`     | Get all tasks             |
| **GET**    | `/tasks/:id` | Get a specific task by ID |
| **PUT**    | `/tasks/:id` | Update an existing task   |
| **DELETE** | `/tasks/:id` | Delete a task by ID       |


## Example Requests(Postman)

**Create Task**
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Understand REST APIs","completed":false}'

**Get All Tasks**
curl -X GET http://localhost:8080/tasks

**Get Task by ID**
curl -X GET http://localhost:8080/tasks/1

**Update Task**
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Master Gin framework","completed":true}'

**Delete Task**
curl -X DELETE http://localhost:8080/tasks/1


## JWT AUTHENTICATION 

**Login and Get Token**
POST /login
{
  "username": "ritik",
  "password": "12345"
}

**Response**
{
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}

**Access Protected Routes**
Add this Header in your requests
Authorization: Bearer <token>

Example:
curl -X GET http://localhost:8080/tasks \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR..."

## Deploying with Ngrok

**Step 1: Download Ngrok**
Go to https://ngrok.com/download
**Step 2: Connect Your Account**
After creating an Ngrok account, get your auth token from the dashboard and run:
ngrok config add-authtoken <your_token_here>
**Step 3: Expose Your Local API**
ngrok http 8080

Response:
Forwarding  https://abcd-1234.ngrok-free.app -> http://localhost:8080

Now your API is accessible globally at:
https://abcd-1234.ngrok-free.app/tasks

