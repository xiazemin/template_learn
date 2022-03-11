// example6.go
package main

import (
	"fmt"
	"runtime"
)

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	a := 1.21111111
	fmt.Printf("%#T   %#T\n", a, 1.21111111)
	//float64 float64
	fmt.Println(1.21111111 == 1.21111112)
	//false
	fmt.Println(float32(1.21111111) == float32(1.21111112))
	//true
	fmt.Println(equal[float32](1.21111111, 1.21111112))
	/*
		equal
		1.2111111
	*/

	fmt.Println(float64(1.2111111111111111111111) == float64(1.2111111111111111111112))
	//true
	fmt.Println(float64(1.21111111) == float64(1.21111112))
	//false
	fmt.Println(equal[float64](1.21111111, 1.21111112))
	// 1.21111112

	/*
		var value1 complex64 = 3.2 + 12i
		//fmt.Println(Change[complex64](value1))
		//./main.go:60:14: Change[complex64](value1) (no value) used as value
		fmt.Println(Change(real(value1)))
		// ./main.go:62:14: Change(real(value1)) (no value) used as value
		fmt.Println(Change(imag(value1)))
		// ./main.go:63:14: Change(imag(value1)) (no value) used as value
	*/
	//https://studygolang.com/articles/13139
	test()
	fmt.Println(runtime.Version())
	go119()
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

type number interface {
	~int | int32 | int64
}

type numberInterface interface {
	~int | int32 | int64 | interface{}
}

type numberAny interface {
	~int | int32 | int64 | any
}

type Interface1 interface {
	Add(a, b int) int
}

// type numberInterface1 interface {
// 	~int | int32 | int64 | Interface1
// }

//./main.go:97:25: cannot use main.Interface1 in union (main.Interface1 contains methods)

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type StructGen[T any] struct {
	A T
	B *T
}

// type StructGenUnamed[T any] struct {
// 	T
// }

//./main.go:137:2: embedded field type cannot be a (pointer to a) type parameter

type InterfaceGen[T any] interface{}

type InterfaceGenWithFunc[T any] interface {
	Add(a, b T) T
}

// type InterfaceGenWithFuncT[T any] interface {
// 	Add[TF any](a, b TF) TF
// }

/*
./main.go:149:5: interface method must have no type parameters
./main.go:149:19: undefined: TF
*/

// type InterfaceEmbed interface {
// 	InterfaceGenWithFunc
// }

//./main.go:158:2: cannot use generic type InterfaceGenWithFunc[T any] without instantiation

// type InterfaceEmbed[T any] interface {
// 	InterfaceGenWithFunc
// }
//./main.go:164:2: cannot use generic type InterfaceGenWithFunc[T any] without instantiation

// type Foo struct{}

// func (Foo) bar[T any](t T) {}

//./main.go:170:15: syntax error: method must have no type parameters
