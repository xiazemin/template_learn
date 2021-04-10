package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

var t string = `{{ define "wrapper" }}
name:{{ .Name }}
end
age: {{ .Age }}
friends: {{ range .Friends }}
friendName: {{.Name}}
{{end}}
{{ end }}
`

type Person struct {
	Name    string
	Age     int
	Friends []Person
}

func main() {
	tmp, err := template.New("text").Parse(t)
	if err != nil {
		panic(err)
	}
	d, e := json.Marshal(tmp)
	builder := &strings.Builder{}
	tmp.ExecuteTemplate(builder, "wrapper", Person{Name: "name", Age: 123, Friends: []Person{{
		Name: "name1", Age: 1234,
	}}})
	fmt.Println(string(d), e, builder.String())
}

//{"Name":"text","ParseName":"text","Root":{"NodeType":11,"Pos":0,"Nodes":[{"NodeType":0,"Pos":129,"Text":"Cg=="}]}} <nil>
