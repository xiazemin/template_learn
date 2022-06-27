package main

import (
	"errors"
	"fmt"
	"template_learn/error/package1"
	"template_learn/error/package2"
)

// Create a named type for our new error type.
type errorString string

// Implement the error interface
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error.
func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}

	if package1.ErrNamedType == New("EOF") {
		fmt.Println("package1 Named Type Error")
	}

	if package1.ErrStructType == errors.New("EOF") {
		fmt.Println("package1 Struct Type Error")
	}

	if package2.ErrNamedType == New("EOF") {
		fmt.Println("package2 Named Type Error")
	}

	if package2.ErrStructType == errors.New("EOF") {
		fmt.Println("package2 Struct Type Error")
	}

	if errors.Is(package2.ErrStructType, errors.New("EOF")) {
		fmt.Println("package2 Struct Type Error")
	}

	if errors.Is(package2.ErrNamedType, New("EOF")) {
		fmt.Println("package2 Struct Type Error")
	}
}

//https://www.cnblogs.com/zhangpengfei5945/p/14470279.html
