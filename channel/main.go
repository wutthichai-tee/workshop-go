package main

import "fmt"

func main() {
	var chnobuff chan int
	chnobuff = make(chan int)

	go func() {
		chnobuff <- 9
	}()

	fmt.Println("Channel without buffer, value received:", <-chnobuff)
}
