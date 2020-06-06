package main

// TodoService is responsible for todos
type TodoService struct {
	lastID int
	todos  []Todo
}

// NewTodoService creates new service
func NewTodoService(todos []Todo) *TodoService {
	if len(todos) == 0 {
		return &TodoService{lastID: 0, todos: todos}
	}

	return &TodoService{lastID: todos[len(todos)-1].ID, todos: todos}
}

// ListTodos |
func (ts *TodoService) ListTodos() []Todo {
	return ts.todos
}

// CreateTodo |
func (ts *TodoService) CreateTodo(task string) {
	todo := Todo{ID: ts.lastID + 1, Task: task, IsCompleted: false}
	ts.lastID++
	ts.todos = append(ts.todos, todo)
}

// UpdateTodo |
func (ts *TodoService) UpdateTodo(todo Todo) {
	for i, t := range ts.todos {
		if t.ID == todo.ID {
			ts.todos[i] = todo
			return
		}
	}
}

// CompleteTodo |
func (ts *TodoService) CompleteTodo(id int) {
	for i, t := range ts.todos {
		if t.ID == id {
			ts.todos[i] = Todo{
				ID:          t.ID,
				Task:        t.Task,
				IsCompleted: true,
			}
			return
		}
	}
}

// DeleteTodo |
func (ts *TodoService) DeleteTodo(id int) {
	newTodos := []Todo{}

	for _, t := range ts.todos {
		if t.ID != id {
			newTodos = append(newTodos, t)
		}
	}

	ts.todos = newTodos
}
