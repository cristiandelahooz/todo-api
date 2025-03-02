package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: "1", Title: "Aprender Go", Status: "pendiente"},
	{ID: "2", Title: "Make an api", Status: "In process"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.Bind(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func updateTodoByID(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		return
	}
	for i, t := range todos {
		if t.ID == id {
			todos[i].Status = updatedTodo.Status
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func deleteTodoByID(c *gin.Context) {
	id := c.Param("id")

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task was eliminated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})

}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/todos")
	})
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", updateTodoByID)
	router.DELETE("/todos/:id", deleteTodoByID)

	router.Run("localhost:8080")
}
