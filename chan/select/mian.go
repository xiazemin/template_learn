package main

import "fmt"

func main() {
	c := make(chan int, 10)
	for i := 0; i < 5; i++ {
		c <- i
	}
	select {
	case v := <-c:
		fmt.Println(v)
	default:
		fmt.Println("default")
	}
	fmt.Println("for")
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		default:
			fmt.Println("default")
		}
	}

}
