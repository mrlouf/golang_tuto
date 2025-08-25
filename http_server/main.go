package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func getTask() {

}

func gracefulShutdown() {

	log.Println("Graceful shutdown of the server...")

}

func main() {

	log.Println("Starting HTTP server on :8080...")

/* 	fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs) */

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