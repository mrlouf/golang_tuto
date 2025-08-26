package api

import (
	"log"
	"encoding/json"
	"net/http"
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