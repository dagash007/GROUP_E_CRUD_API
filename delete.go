// delete.go
package main

import (
	"net/http"
	"strconv"
	"strings"
)

// deleteTask handles DELETE /tasks/{id}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	// Search for the task to delete
	for i, task := range tasks {
		if task.ID == id {
			// Remove the task from the slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			// Respond with no content
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	// If task not found
	http.Error(w, "Task Not Found", http.StatusNotFound)
}
