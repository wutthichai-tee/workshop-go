package main

import (
	"fmt"
	"time"
)

var i int

func main() {
	fmt.Println("Hello, World!")

	start := time.Now()

	for range 5 {
		go slow(func() {
			fmt.Println(i)
		})
	}

	for {
		if i >= 4 {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func slow(fn func()) {
	time.Sleep(100 * time.Millisecond)
	i++
	fn()
}
