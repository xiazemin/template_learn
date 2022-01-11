package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
			time.Sleep(time.Second)
		}
	}()
	for i := 1; i < 10; i++ {
		select {
		case c <- i:
			fmt.Println("c <- i")
		default:
			fmt.Println("default")
		}

	}
	time.Sleep(3 * time.Second)
}
