package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func run(file string, ch chan int) {
	
	data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println(err)
    }
	ch <- len(data)
}

func main() {
    argsWithoutProg := os.Args[1:]

	ch := make(chan int)

	fmt.Println(len(argsWithoutProg), "files to read")

	var x int = 0

	for _, filename := range argsWithoutProg {
		go run(filename, ch)
		var tmp = <- ch
		fmt.Println(tmp)
		x = x + tmp
	}

	fmt.Println(x)
}
