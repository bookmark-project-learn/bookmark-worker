package main

import (
	"fmt"
	"time"
)

func receive(str <-chan string) {
	for {
		fmt.Print("start ")

		v, ok := <-str
		fmt.Print("start2 ")
		if !ok {
			fmt.Print("closed")
			return
		} else {
			fmt.Print("open ")
		}
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	str := make(chan string, 5)
	go receive(str)
	str <- "a"
	str <- "b"
	str <- "c"
	time.Sleep(10 * time.Second)

}
