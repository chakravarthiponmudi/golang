package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println(<-c)
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int, 100)
	quit := make(chan int)
	go func() {
		time.Sleep(5 * time.Microsecond)
		quit <- 0
	}()
	fibonacci(c, quit)
}
