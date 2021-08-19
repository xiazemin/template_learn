//方法表达式
//https://www.cnblogs.com/phpper/p/12370086.html
package main

import "fmt"

type My struct {
}

func (m *My) GetInfo() {
	fmt.Println("my info")
}

var MyGetInfo = (*My).GetInfo

func main() {
	m := &My{}
	MyGetInfo(m)
}
