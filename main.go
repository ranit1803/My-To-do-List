package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initDB() {
	var err error
	connStr := "user=postgres password=ranit1803 dbname=todo-app sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

// Task struct
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Get all tasks
func getTasks(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, completed FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

// Add a new task
func addTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Println("JSON Binding Error:", err) // Debug JSON errors
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, task.Title, task.Completed).Scan(&task.ID)

	if err != nil {
		log.Println("Database Insert Error:", err) // Debug SQL errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Update task status
func completeTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("UPDATE tasks SET completed = TRUE WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task completed"})
}

// Delete a task
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func main() {
	initDB()
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/tasks", getTasks)
	router.POST("/tasks", addTask)
	router.PUT("/tasks/:id", completeTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run(":8080")
}
