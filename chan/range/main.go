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
		}
	}()
	for i := 1; i < 10; i++ {
		c <- i
		time.Sleep(time.Second)
	}
}
