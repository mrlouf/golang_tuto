package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"

	"tutogo/mod/http_server/api"
	"tutogo/mod/http_server/utils"
)

func main() {

	store := &api.TaskStore{
		Tasks:	 []api.Task{},
		Counter: 0,
	}

	log.Println("Starting HTTP server on :8080...")
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        api.TaskHandler(w, r, store)
    })
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
	    api.TaskHandler(w, r, store)
	})

	// Listen for signals to shutdown the server
    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        utils.GracefulShutdown()
        os.Exit(0)
    }()
	
    err := http.ListenAndServe(":8080", nil)

	fmt.Println("ListenAndServe is blocking, meaning this should only print in case of an error at start")
    if err != nil {
        log.Fatal(err)
    }

}