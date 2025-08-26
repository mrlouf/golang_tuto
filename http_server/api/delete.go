package api

import (
	"log"
	"strings"
	"strconv"
	"encoding/json"
	"net/http"
)

func HttpHandlerDelete(w http.ResponseWriter, r *http.Request, store *TaskStore) {

	log.Println("DELETE request received")

	if len(store.Tasks) == 0 {
		http.Error(w, "Task list empty, no task deleted", http.StatusBadRequest)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/tasks")

	if path == "" || path == "/" {
		log.Println("No task ID in URL")
		return
	} else {
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		log.Printf("Task ID in URL: %d\n", id)
		for idx, t := range store.Tasks {
			if t.ID == id {
				store.Tasks = append(store.Tasks[:idx], store.Tasks[idx+1:]...)
				break
			}
		}
	}

	response := APIResponse{
		Status:  http.StatusOK,
		Message: "Successfully deleted task",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}