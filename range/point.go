package main

import "fmt"

type a struct {
	V int
}

func main() {
	v := make(map[*a]int)
	for i := 0; i < 5; i++ {
		v[&a{V: i}] = i
	}

	v1 := make(map[*a]int)
	for k, vv := range v {
		fmt.Println(k.V, vv)
		v1[k] = vv
	}

	for k, vv := range v {
		if vvv, ok := v1[k]; ok {
			fmt.Println("in", k.V, vv, vvv)
		}
		fmt.Println(k.V, vv)
	}
}
