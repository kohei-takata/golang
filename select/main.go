package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	// d := make(chan int)

	go func() {
		for {
			a <- 0
		}
	}()

	go func() {
		for {
			b <- 0
		}
	}()

	go func() {
		for {
			c <- 0
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			fmt.Println("a")
		case <-b:
			fmt.Println("b")
		case <-c:
			fmt.Println("c")
		}
	}
}
