package api

import (
	"net/http"
	"log"
	"encoding/json"
	"strings"
	"strconv"
)

func HttpHandlerGet(w http.ResponseWriter, r *http.Request, store *TaskStore) {
	
		log.Println("GET request received")

	    response := APIResponse{
        	Status:  http.StatusOK,
   	    	Data:    store.Tasks,
       		Message: "Successfully retrieved tasks",
	    }

  		w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(response)

}

func TaskHandler(w http.ResponseWriter, r *http.Request, store *TaskStore) {
    log.Println("/tasks endpoint reached, handling request...")
	
	switch r.Method {
	
	case http.MethodGet:
		HttpHandlerGet(w, r, store)

	case http.MethodPost:
		log.Println("POST request received")

		store.Tasks = append(store.Tasks, Task{store.Counter, "buy milk", false})
		store.Counter++

	case http.MethodPut:
		log.Println("PUT request received")

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
			for idx := range store.Tasks {
				if store.Tasks[idx].ID == id {
					store.Tasks[idx].Done = true
					break
				}
			}
		}

	case http.MethodDelete:
		log.Println("DELETE request received")

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

	default:
		http.Error(w, "Invalid Method", http.StatusNotImplemented)
		log.Printf("Error 501: Not Implemented: %s\n", r.Method)
	}
}