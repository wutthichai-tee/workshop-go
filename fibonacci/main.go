package main

import "fmt"

func main() {
	ch := make(chan int)
	qch := make(chan struct{})

	go fibonacci(ch, qch)

	for range 10 {
		fmt.Println(<-ch)
	}

	qch <- struct{}{} // Signal to exit the goroutine

	<-qch // Wait for the goroutine to acknowledge the exit signal
	fmt.Println("Bye")
}

func fibonacci(ch chan int, qch chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
			fmt.Println("A =", a, "B =", b)
		case <-qch:
			fmt.Println("Exiting goroutine in fibonacci")
			qch <- struct{}{}
			return
		}
	}
}
