package main

import "fmt"

type Sortable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

func bubbleSortSlice(array Sortable) {
	for i := 0; i < array.Len(); i++ {
		for j := 0; j < array.Len()-i-1; j++ {
			if array.Less(j+1, j) {
				array.Swap(j, j+1)
			}
		}
	}
}

func bubbleSort(array []interface{}) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			switch array[j].(type) {
			case int:
				arrayIntj1 := array[j+1].(int)
				arrayIntj := array[j].(int)
				if arrayIntj1 < arrayIntj {
					array[j+1], array[j] = array[j], array[j+1]
				}
			}
		}
	}
}

func equal[T comparable](a, b T) T {
	if a == b {
		fmt.Println("equal")
		return a
	}
	return b
}

/*
func compare[T comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func add[T comparable](a, b T) T {
	return a + b
}
*/

//https://tip.golang.org/doc/go1.18
// ./sort.go:42:5: invalid operation: cannot compare a < b (operator < not defined on T)
// ./sort.go:49:9: invalid operation: operator + not defined on a (variable of type T constrained by comparable)
