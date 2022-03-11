package main

import (
	"constraints"
	"fmt"
	"strconv"
)

func call() {
	// 调用泛型函数
	m := min[int](2, 3)
	fmt.Println(m)
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Tree[T interface{}] struct {
	left, right *Tree[T]
	data        T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	return nil
}

var stringTree Tree[string]

type Integer int
type Float = float64
type Ordered interface {
	Integer | Float | ~string
}

// type Number interface{
// 	int | int32 | int64
// }

type AnyString interface {
	~string
	Add(a, b int) int
}
type MyString string

func methodCall() {
	// var m MethodCall
	// fmt.Println(m.add(1, 2))
	//./gen.go:48:8: cannot use generic type MethodCall[T int|int32, TA any] without instantiation

	var m MethodCall[int, int32]
	fmt.Println(m.add(1, 2))
}

type MethodCall[T int | int32, TA any] struct {
	v T
}

func (m MethodCall[T, TA]) add(a, b T) T {
	return a + b
}

// func (m MethodCall[T, TA]) addAny(a, b TA) TA {
// 	return a + b
// }
//./gen.go:59:9: invalid operation: operator + not defined on a (variable of type TA constrained by any)

// 再比如，我们定义一个新的类型限制叫customConstraint，用于限定底层类型为int并且实现了String() string方法的所有类型。下面的customInt就满足这个type constraint。

type customConstraint interface {
	~int
	String() string
}

type customInt int

func (i customInt) String() string {
	return strconv.Itoa(int(i))
}

// 类型限制有2个作用：

// 用于约定有效的类型实参，不满足类型限制的类型实参会被编译器报错。
// 如果类型限制里的所有类型都支持某个操作，那在代码里，对应的类型参数就可以使用这个操作。
//https://segmentfault.com/a/1190000041174189

// SliceConstraint 匹配所有类型为 T 的切片，但 T 的类型需要是使用的时候指定！
type SliceConstraint[T any] interface {
	[]T
}

//https://taoshu.in/go/generics/design.html
//https://www.cnblogs.com/ink19/p/go_generic_programming.html
//https://juejin.cn/post/7042949225138782244

type MyStruct[P any] struct {
	v  P
	vf func(p2 P)
}

func (m MyStruct[p]) m(p1 p) {
	// fmt.Println(p1 + 1)
	//./gen.go:103:19: cannot convert 1 (untyped int constant) to p
	fmt.Println(p1)
}

func support() {
	var x MyStruct[any]
	x.m(12)
	fmt.Println(x.v)
	x.vf(23)
	//(*MyStruct).m(&x, 12)
	//./gen.go:115:4: cannot use generic type MyStruct without instantiation
	// (*MyStruct[int]).m(&x, 12)
	//./gen.go:117:21: cannot use &x (value of type *MyStruct[any]) as type *MyStruct[int] in argument to (*MyStruct[int]).m
	var xx MyStruct[int]
	(*MyStruct[int]).m(&xx, 12)
	//goLang 之 type Method Value 和Method Expressions
	//https://studygolang.com/articles/12353?fr=sidebar
	//https://studygolang.com/articles/03806
	var xxx MyStruct[int]
	xxx.m(xx.v)

	// xxxx := MyStruct{
	// 	v: 1234,
	// }
	//./gen.go:127:10: cannot use generic type MyStruct[P any] without instantiation
	xxxx := MyStruct[int]{
		v: 1234,
	}
	xxxx.m(12)

	var s MyStruct[Student]
	s.v.Name = "13224"
	fmt.Println(s.v.Name)
}

type Student struct {
	Name string
}
type Teacher struct {
	Name string
}

func test() {

	var st Student
	st.Name = "hhh"
	Show[Student](st)
	// ShowStudent[Student](st)
	//./gen.go:132:14: Student does not implement Student (Student is not an interface)

	AddOne(1, 2)
}

func ShowStudent[T Student](s T) {
	fmt.Println(s.Name)
}

type MyStudent Student

// func ShowStudentType[T ~Student](s T) {
// 	fmt.Println(s.Name)
// }

//./gen.go:142:24: invalid use of ~ (underlying type of Student is struct{Name string})

func Show[T Student | Teacher](s T) {
	fmt.Println(s.Name)
}

//method expressions

func (s Student) Say() {
	fmt.Println("Student")
}

func (t Teacher) Say() {
	fmt.Println("Teacher")
}

// func Say[T Student | Teacher](p T) {
// 	p.Say()
// }
//./gen.go:183:4: p.Say undefined (type T has no field or method Say)
func Say[T interface{ Say() }](p T) {
	p.Say()
}

func AddOne[T int](a, b T) T {
	return a + b
}
