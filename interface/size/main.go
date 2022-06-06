package main

import (
	"fmt"
	"unsafe"
)

type I interface {
	Get() int
	Put(int)
	A() int //可以自由添加，只为检验是否增加后，会改变占用字节

}

func main() {
	var i I
	var j interface{}
	fmt.Println(unsafe.Sizeof(i), unsafe.Sizeof(j))
	//16 16
}

//eface和iface都占16字节，eface只有2个字段，因为它代表的是没有方法的接口，只需要存储被赋值对象的类型和数据即可，正好对应到这里的_type 和 data字段。iface代表含有方法的接口，定义里面的 data字段也是表示被存储对象的值，注意这里的值是原始值的一个拷贝，如果原始值是一个值类型，这里的data是执行的数据时原始数据的一个副本。
