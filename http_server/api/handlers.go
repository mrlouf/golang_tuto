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

func HttpHandlerPost(w http.ResponseWriter, r *http.Request, store *TaskStore) {
	log.Println("POST request received")

	var input = ""
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil || input == "" {
        input = "Untitled Task"
    }

    newTask := Task{
        ID:    store.Counter,
        Title: input,
        Done:  false,
    }
    store.Tasks = append(store.Tasks, newTask)
    store.Counter++

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newTask)
}

func TaskHandler(w http.ResponseWriter, r *http.Request, store *TaskStore) {

    log.Println("/tasks endpoint reached, handling request...")
	
	switch r.Method {
	
	case http.MethodGet:
		HttpHandlerGet(w, r, store)

	case http.MethodPost:
		HttpHandlerPost(w, r, store)

	case http.MethodPut:
		HttpHandlerPut(w, r, store)

	case http.MethodDelete:
		HttpHandlerDelete(w, r, store)

	default:
		http.Error(w, "Invalid Method", http.StatusNotImplemented)
		log.Printf("Error 501: Not Implemented: %s\n", r.Method)
	}
}