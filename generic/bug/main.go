package main

func main() {
	testBefore()
	testAfter()
}
func testBefore() {
	var st Student
	Show1[Student](st)
	Show2[Student](st)
}
func Show2[T Student](s T)           {}
func Show1[T Student | Teacher](s T) {}

type Student struct{}
type Teacher struct{}

func testAfter() {
	var st Student
	Show1[Student](st)
	Show2[Student](st)
}
