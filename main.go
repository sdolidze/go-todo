package main

import (
	"log"
	"net/http"
)

func main() {
	todoController := NewTodoController()

	http.HandleFunc("/", todoController.Index)
	http.HandleFunc("/todos", todoController.ListTodos)
	http.HandleFunc("/todos/create", todoController.CreateTodo)
	http.HandleFunc("/todos/delete", todoController.DeleteTodo)
	http.HandleFunc("/todos/complete", todoController.CompleteTodo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
