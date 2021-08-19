package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	b := 2
	fmt.Println(a + b)
	c := make(chan struct{})
	c0 := make(chan struct{}, 0)
	c1 := make(chan struct{}, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			v, ok := <-c
			if ok {
				fmt.Println("ok", v)
			} else {
				fmt.Println("not ok", v)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v, ok0 := <-c0
			if ok0 {
				fmt.Println("ok0", v)
			} else {
				fmt.Println("not ok0", v)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v, ok1 := <-c1
			if ok1 {
				fmt.Println("ok1", v)
			} else {
				fmt.Println("not ok1", v)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v, ok2 := <-c2
			if ok2 {
				fmt.Println("ok2", v)
			} else {
				fmt.Println("not ok2", v)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v, ok3 := <-c3
			if ok3 {
				fmt.Println("ok3", v)
			} else {
				fmt.Println("not ok3", v)
			}
		}
	}()

	tick := time.NewTicker(1 * time.Second)
	select {
	case <-tick.C:
		close(c)
		close(c0)
		close(c1)
		close(c2)
		c3 <- 1
		close(c3)
		//case c<- struct{}{} :
		fmt.Println("select")
	}
	time.Sleep(1 * time.Millisecond)
	//c<-struct{}{}
	/**
	  3
	select
	ok3 1
	not ok3 0
	not ok3 0
	not ok3 0
	not ok3 0
	not ok3 0
	not ok3 0
	not ok3 0
	not ok3 0
	not ok3 0
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok0 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok1 {}
	not ok2 0
	not ok2 0
	not ok2 0
	not ok2 0
	not ok2 0
	not ok2 0
	not ok2 0
	not ok2 0
	not ok2 0
	not ok2 0
	not ok {}
	not ok {}
	not ok {}
	not ok {}
	not ok {}
	not ok {}
	not ok {}
	not ok {}
	not ok {}
	not ok {}
	*/
}

//close 后chan 不会阻塞，但是可以从里面取出0值
