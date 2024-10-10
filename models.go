// models.go
package main

import (
	"sync"
)

// Task represents a task with ID, Title, Description, and Status
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

// In-memory storage for tasks
var (
	tasks  = []Task{}
	nextID = 1
	mu     sync.Mutex // Mutex to handle concurrent access
)
