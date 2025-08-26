package api

import (
	"net/http"
	"log"
	"encoding/json"
)

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