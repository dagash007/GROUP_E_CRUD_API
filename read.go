// read.go
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// getTasks handles GET /tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// getTaskByID handles GET /tasks/{id}
func getTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Search for the task
	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	// If task not found
	http.Error(w, "Task Not Found", http.StatusNotFound)
}
