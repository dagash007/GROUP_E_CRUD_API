// update.go
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// updateTask handles PUT /tasks/{id}
func updateTask(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	// Decode the JSON request body into the Task struct
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	// Search for the task to update
	for i, task := range tasks {
		if task.ID == id {
			// Update fields if provided
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				if updatedTask.Status == "pending" || updatedTask.Status == "completed" {
					tasks[i].Status = updatedTask.Status
				} else {
					http.Error(w, "Invalid Status Value", http.StatusBadRequest)
					return
				}
			}
			// Set response header to JSON
			w.Header().Set("Content-Type", "application/json")
			// Respond with the updated task
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}

	// If task not found
	http.Error(w, "Task Not Found", http.StatusNotFound)
}
