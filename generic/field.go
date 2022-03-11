package main

import (
	"fmt"
	"runtime"
)

func go1119() {
	fmt.Println(runtime.Version())
	var st Student
	st.Name = "hhhx"
	Show1[Student](st) //this can work
	Show2[Student](st)

	Show[Student](st)
	// ShowStudent[Student](st)
}

func Show2[T Student](s T) {
	// fmt.Println(s.Name)
}

func Show1[T Student | Teacher](s T) {
	// fmt.Println(s.Name)
}

func go119() {
	fmt.Println(runtime.Version())
	var st Student
	st.Name = "hhhx"
	Show1[Student](st) //this can work
	Show2[Student](st)

	Show[Student](st)
	// ShowStudent[Student](st)
}
