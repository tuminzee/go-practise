package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- 42
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Hello, world!"
	}()

	select {
	case msg := <-c1:
		fmt.Println("Received from c1:", msg)
	case msg := <-c2:
		fmt.Println("Received from c2:", msg)
	default:
		fmt.Println("Timeout: No data received within 3 seconds.")
	}
}
