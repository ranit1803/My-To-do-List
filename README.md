# To-Do List Application

This is a simple To-Do List application built using Gin framework in Go, PostgreSQL for the database, and a frontend built with HTML, CSS, and JavaScript.

## Features

-> Add Tasks: Users can add tasks to their to-do list.
-> Delete Tasks: Users can delete tasks from their list.
-> Complete Tasks: Users can mark tasks as completed.
-> Task Filtering: Filter tasks by all, completed, or incomplete tasks.
-> Frontend: A clean and simple UI to manage tasks.

## Technologies Used

-> Backend: 
  -> Go: Backend API built using the Gin framework.
  -> PostgreSQL: A relational database to store tasks.
-> Frontend: 
  -> HTML/CSS: For structuring and styling the web page.
  -> JavaScript: For interacting with the backend API and handling user actions like adding, completing, and deleting tasks.
-> Version Control: Git, hosted on GitHub.

## Setup Instructions
### Prerequisites

-> Go 1.17+ installed on your machine.
-> PostgreSQL installed and running.
-> Git installed on your machine.
-> A text editor or IDE (e.g., VSCode, IntelliJ, etc.).
Step 1: Clone the repository:
        ```bash
        git clone https://github.com/ranit1803/My-To-do-List.git
        cd My-To-do-List
Step 2: go mod tidy
Step 3: Setup Postgres SQL Database
        CREATE DATABASE todo-app;
        CREATE TABLE tasks (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            completed BOOLEAN DEFAULT FALSE
        );
Step 4: In the Go backend code (main.go), update the database connection string with your PostgreSQL credentials.
Step 5: go run main.go 
        Open index.html file
