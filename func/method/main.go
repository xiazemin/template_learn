package main

import "fmt"

type St1 struct {
	//A int
}

/*
0x100442fc8
0x100442fc8
0x100442fc8
0x100442fc8
0x100442fc8
0x100442fc8
0x100442fc8
0x100442fc8
*/
type St struct {
	A int
}

/*
0x14000126008
0x14000126008
0x14000126020
0x14000126028
0x14000126030
0x14000126030
0x14000126038
0x14000126040
*/

func (this *St) func0() {
	fmt.Printf("%p\n", this)
}
func (this *St) func1() {
	fmt.Printf("%p\n", this)
}

//调用一次copy一次
func (this St) func2() {
	fmt.Printf("%p\n", &this)
}
func (this St) func3() {
	fmt.Printf("%p\n", &this)
}

func main() {
	var st1, st11 St1
	st3 := &St1{}
	st4 := &St1{}
	st5 := St1{}
	st6 := St1{}
	fmt.Printf("%p\n%p\n%p\n%p\n%p\n%p\n\n", &st1, &st11, st3, st4, &st5, &st6)
	/*
		0x10474efc8
		0x10474efc8
		0x10474efc8
		0x10474efc8
		0x10474efc8
		0x10474efc8
	*/

	var st St

	st.func0()
	st.func1()
	st.func2()
	st.func3()
	st2 := st
	st2.func0()
	st2.func1()
	st2.func2()
	st2.func3()
}
