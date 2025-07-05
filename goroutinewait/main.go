package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Main function started executing")

	wg := &sync.WaitGroup{}

	wg.Add(3)
	go works("go", wg)
	go works("rust", wg)
	go works("angular", wg)
	wg.Wait()

	fmt.Println("Main function finished executing")
}

func works(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	b := sha256.Sum256([]byte(s))
	fmt.Println(base64.StdEncoding.EncodeToString(b[:]))
}
