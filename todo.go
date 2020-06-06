package main

// Todo is a unit of work
type Todo struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	IsCompleted bool   `json:"isCompleted"`
}
