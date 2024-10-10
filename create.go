// create.go
package main

import (
	"encoding/json"
	"net/http"
)

// createTask handles POST /tasks
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Decode the JSON request body into the Task struct
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure title and description are provided
	if task.Title == "" || task.Description == "" {
		http.Error(w, "Title and Description are required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	task.ID = nextID
	nextID++
	task.Status = "pending" // Default status
	tasks = append(tasks, task)
	mu.Unlock()

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	// Respond with the created task
	json.NewEncoder(w).Encode(task)
}
