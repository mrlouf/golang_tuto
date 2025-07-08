package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func run(file string, ch chan int) int {
	
	data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println(err)
    }

	return len(data)
}

func main() {
	ch := make(chan int, len(os.Args))

	for _, filename := range os.Args {
		go run(filename, ch)
	}

	x := <- ch
	fmt.Println(x)
}
