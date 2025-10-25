package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"Completed"`
}

// getTodos return all todos as json
func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// getTodosById return todo whose id matches the id setn by client
func getTodoById(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// getTodoByTitle return todo that match title sent by client (as json)
func getTodoByTitle(c *gin.Context) {
	title := c.Param("title")

	for _, todo := range todos {
		if todo.Title == title {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// getCompletedTodos return completed todos as json
func getCompletedTodos(c *gin.Context) {
	var completed_todos []Todo

	for _, todo := range todos {
		if todo.Completed {
			completed_todos = append(completed_todos, todo)
		}
	}

	c.JSON(http.StatusOK, completed_todos)
}

// postTodos add todo
func postTodos(c *gin.Context) {
	var newTodo Todo

	err := c.BindJSON(&newTodo)

	if err != nil {
		return
	}

	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	var updatedTodo Todo

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			todo.Title = updatedTodo.Title
			todo.Description = updatedTodo.Description
			todo.Completed = updatedTodo.Completed

			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// putCompleted update the status of todo to completed
func putCompleted(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			todo.Completed = true

			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	var newTodos []Todo

	for _, todo := range todos {
		if todo.ID == id {
			continue
		}

		newTodos = append(newTodos, todo)
	}

	todos = newTodos
	c.Status(http.StatusNoContent)
}

var todos = []Todo{
	{"1", "Learn Golang", "Understand the basics of golang, Build application", false},
	{"2", "Bath", "Get a shower", false},
	{"3", "Exercise", "Do 2hr exercise focusing on stamina and endurance", false},
	{"4", "Hacknight", "Attend and build staffs during hacknight", true},
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.GET("/todos/title/:title", getTodoByTitle) //get todo with specific title
	router.GET("/todos/completed", getCompletedTodos) //get completed todos
	router.POST("/todos", postTodos)                  //adding new todo
	router.PUT("/todos/:id", updateTodo)              //update details of todo ie title, description etc
	router.PUT("/todos/completed/:id", putCompleted)  //change todo status to completed !better if it was a patch
	router.DELETE("/todos/:id", deleteTodo)           //delete todo

	router.Run("localhost:8000")
}
