package main

import (
	"fmt"
	"regexp"
)

var (
	emailRex, _  = regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	mobileRex, _ = regexp.Compile("1[3-9][0-9]{9}$")
)

func main() {
	fmt.Println(emailRex.Match([]byte("autoTest_company_1@shimo.im")))
}
