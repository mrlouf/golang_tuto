package api

import (
	"log"
	"strings"
	"encoding/json"
	"net/http"
	"strconv"
)

func HttpHandlerPut(w http.ResponseWriter, r *http.Request, store *TaskStore) {

	log.Println("PUT request received")

	path := strings.TrimPrefix(r.URL.Path, "/tasks")
	if path == "" || path == "/" {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	} else {
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		log.Printf("Task ID in URL: %d\n", id)
		for idx := range store.Tasks {
			if store.Tasks[idx].ID == id {
				store.Tasks[idx].Done = true
				break
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.Tasks[id])
	}
}