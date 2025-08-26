package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func run(file string, ch chan int) {
	
	data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println(err)
		ch <- 0
		return
    }
	words := strings.Fields(string(data))
	ch <- len(words)
}

func main() {
    argsWithoutProg := os.Args[1:]

	ch := make(chan int)

	fmt.Println(len(argsWithoutProg), "files to read")

	for _, filename := range argsWithoutProg {
		go run(filename, ch)
	}

	total := 0

	for i := 0; i < len(argsWithoutProg); i++ {
		total += <-ch
	}

	fmt.Println("Total words read:", total)
}
