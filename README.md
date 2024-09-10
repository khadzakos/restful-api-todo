# restful-api-todo
## RESTful API for To-Do List (Golang)

## Overview

This project is a RESTful API for a to-do list application built using Go. The API allows users to create, read, update, and delete (CRUD) tasks.

## Features
**CRUD Operations:**

	•	Create: Add a new task.
	•	Read: View all tasks or a specific task.
	•	Update: Modify task details.
	•	Delete: Remove tasks.
 ## Project Structure
 ````
.
├── database
│   ├── models
│   │   └── task.go
│   └── database.go
├── docs
│   └── openapi.yml
├── handlers
│   └── task_handlers.go
├── main.go
````
## Installation
1.	**Clone the repository:**
````
git clone https://github.com/khadzakos/restful-api-todo.git
````
2. **Run Docker Compose:**
````
docker-compose up --build
````

## API Endpoints
	•	GET /tasks: Retrieve all tasks.
	•	POST /tasks: Create a new task.
	•	PUT /tasks/{id}: Update a task by ID.
	•	DELETE /tasks/{id}: Delete a task by ID.
 
## Technologies Used
-	Go (Golang)
-	PostgreSQL (Database)
-	Docker

## Libraries 
- Gin
- Gorm
