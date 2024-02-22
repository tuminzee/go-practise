package main

import (
	"fmt"
)

func main() {
	const v = "hello world"
	c := make(chan int, 4)

	go func(c chan int) {
		for i := 1; i <= 10; i++ {
			fmt.Printf("SEND BEFORE %d \n", i)
			c <- i
			fmt.Printf("SEND AFTER %d \n", i)
		}
		close(c)
	}(c)

	//time.Sleep(time.Second * 2)

	for v := range c {
		fmt.Printf("RECEIVE %d \n", v)

	}

}
