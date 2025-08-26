package api

import (
	"net/http"
	"log"
)

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