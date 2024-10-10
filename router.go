// router.go
package main

import (
	"net/http"
)

// router sets up the routes and their handlers
func router() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getTasks(w, r)
		case "POST":
			createTask(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle routes with ID (e.g., /tasks/1)
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getTaskByID(w, r)
		case "PUT":
			updateTask(w, r)
		case "DELETE":
			deleteTask(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
