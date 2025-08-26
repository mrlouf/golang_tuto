package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"strconv"
	"encoding/json"
)

type Task struct {
	ID int
	Title string
	Done bool
}

var taskId = 0

func gracefulShutdown() {

	log.Println("Graceful shutdown of the server...")

}

func taskHandler(w http.ResponseWriter, r *http.Request, tasks *[]Task) {
    log.Println("/tasks endpoint reached, handling request...")
	
	switch r.Method {
	
	case http.MethodGet:
		log.Println("GET request received")

		for _, task := range *tasks {

			b, err := json.Marshal(task)
			if err != nil {
				log.Fatalf("Unable to marshal due to %s\n", err)
			}

			fmt.Println(string(b))
		}

	case http.MethodPost:
		log.Println("POST request received")

		*tasks = append(*tasks, Task{taskId, "buy milk", false})
		taskId++

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
			for idx := range *tasks {
				if (*tasks)[idx].ID == id {
					(*tasks)[idx].Done = true
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
			for idx, t := range *tasks {
				if t.ID == id {
					*tasks = append((*tasks)[:idx], (*tasks)[idx+1:]...)
					break
				}
			}
		}

	default:
		http.Error(w, "Invalid Method", http.StatusNotImplemented)
		log.Printf("Error 501: Not Implemented: %s\n", r.Method)
	}
}

func main() {

	tasks := []Task{}

	log.Println("Starting HTTP server on :8080...")
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        taskHandler(w, r, &tasks)
    })
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
	    taskHandler(w, r, &tasks)
	})

	// Listen for signals to shutdown the server
    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        gracefulShutdown()
        os.Exit(0)
    }()
	
    err := http.ListenAndServe(":8080", nil)

	fmt.Println("ListenAndServe is blocking, meaning this should only print in case of an error at start")
    if err != nil {
        log.Fatal(err)
    }

}