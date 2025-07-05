package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	i   int
	mux sync.Mutex
}

var asset data

func main() {

	// Command to check race conditions
	// go run -race main.go
	fmt.Println("Hello, World!")

	start := time.Now()

	for range 5 {
		go slow(func() {
			asset.mux.Lock()
			fmt.Println(asset.i)
			asset.mux.Unlock()
		})
	}

	for {
		asset.mux.Lock()
		if asset.i >= 4 {
			break
		}
		asset.mux.Unlock()
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func slow(fn func()) {
	time.Sleep(100 * time.Millisecond)
	asset.mux.Lock()
	asset.i++
	asset.mux.Unlock()
	fn()
}
