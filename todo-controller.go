package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// TodoController |
type TodoController struct {
	todoService *TodoService
}

// NewTodoController |
func NewTodoController() *TodoController {
	todoService := NewTodoService([]Todo{
		{ID: 1, Task: "Learn Go", IsCompleted: false},
		{ID: 2, Task: "Learn Docker", IsCompleted: true},
		{ID: 3, Task: "Learn Kubernetes", IsCompleted: false},
	})

	return &TodoController{todoService: todoService}
}

// Index |
func (*TodoController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Todo Server is Go!")
}

// ListTodos |
func (tc *TodoController) ListTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tc.todoService.ListTodos())
}

// CreateTodo |
func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	tasks, ok := r.URL.Query()["task"]

	if !ok || len(tasks[0]) < 1 {
		fmt.Fprintf(w, "Url parameter 'task' is missing")
		log.Println("Url parameter 'task' is missing")
		return
	}

	task := tasks[0]

	tc.todoService.CreateTodo(task)
	fmt.Fprintf(w, "New task created")
}

// DeleteTodo |
func (tc *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]

	if !ok || len(ids[0]) < 1 {
		fmt.Fprintf(w, "Url parameter 'id' is missing")
		log.Println("Url parameter 'id' is missing")
		return
	}

	id, err := strconv.Atoi(ids[0])

	if err != nil {
		fmt.Fprintf(w, "Url parameter 'id' is not a valid number")
		log.Println("Url parameter 'id' is not a valid number")
		return
	}

	tc.todoService.DeleteTodo(id)

	fmt.Fprintf(w, "Task with id %d has been deleted", id)
}

// CompleteTodo |
func (tc *TodoController) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]

	if !ok || len(ids[0]) < 1 {
		fmt.Fprintf(w, "Url parameter 'id' is missing")
		log.Println("Url parameter 'id' is missing")
		return
	}

	id, err := strconv.Atoi(ids[0])

	if err != nil {
		fmt.Fprintf(w, "Url parameter 'id' is not a valid number")
		log.Println("Url parameter 'id' is not a valid number")
		return
	}

	tc.todoService.CompleteTodo(id)

	fmt.Fprintf(w, "Task with id %d has been completed", id)
}
