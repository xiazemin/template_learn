package main
import (
"fmt"
"regexp"
)

func main(){
	emailRex, _     := regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
       fmt.Println(emailRex.Match([]byte("12345678@123456.ip")))
  fmt.Println(emailRex.Match([]byte("12345678@111111111111.ip")))
	emailRex1, _      := regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+)+[a-zA-Z]*))$`)
 fmt.Println(emailRex1.Match([]byte("12345678@123456.ip")))
 fmt.Println(emailRex1.Match([]byte("12345678@123456")))
	emailRex2, _      := regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,})|.*)$`)

fmt.Println(emailRex2.Match([]byte("12345678@123456.ip")))
fmt.Println(emailRex2.Match([]byte("12345678@123456")))
fmt.Println(emailRex2.Match([]byte("12345678123456")))
rex,_:=regexp.Compile(`^.+@.+$`)
fmt.Println(rex.Match([]byte("123456")))
fmt.Println(rex.Match([]byte("12345@123")))
fmt.Println(rex.Match([]byte("@123")))
fmt.Println(rex.Match([]byte("qww@")))
}
