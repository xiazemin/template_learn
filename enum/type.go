package main

import "fmt"

type Typ1 int32

const (
	Typ1_One Typ1 = 1
	Typ1_Two Typ1 = 2
)

type Typ2 int32

const (
	Typ2_Two   Typ2 = 2
	Typ2_Three Typ2 = 3
)

func main() {
	t2 := Typ2_Two
	t3 := Typ2_Three
	fmt.Println(Typ1(t2), Typ1(t3))
	fmt.Printf("%T", Typ1(t3))
}
