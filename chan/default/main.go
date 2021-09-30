package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)
	c <- 1
	select {
	case v := <-c:
		fmt.Println(v)
	default:
		fmt.Println("default")
	}
	c <- 1
	time.Sleep(5000)
}
