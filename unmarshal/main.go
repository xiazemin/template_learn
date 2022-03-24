package main

import (
	"encoding/json"
	"fmt"
)

type St struct {
	Name string
}

func main() {
	var v *St
	err := json.Unmarshal([]byte(`{"name":"xzm"}`), v)
	fmt.Println(err, v)
	v1 := &St{}
	err1 := json.Unmarshal([]byte(`{"name":"xzm"}`), v1)
	fmt.Println(err1, v1)
}

/*
json: Unmarshal(nil *main.St) <nil>
<nil> &{xzm}
*/
