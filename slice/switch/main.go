package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for _, a := range arr {
		switch a {
		case 1:
			fmt.Println(1)
		case 2:
			break
		case 3, 4, 5:
			fmt.Println(a)
		}
		fmt.Println("aaa:", a)
	}
}
