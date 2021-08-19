package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	tmp, err := template.New("text").Parse(t)
	if err != nil {
		panic(err)
	}
	d, e := json.Marshal(tmp)
	fmt.Println(string(d), e)
}

//{"Name":"text","ParseName":"text","Root":{"NodeType":11,"Pos":0,"Nodes":[{"NodeType":0,"Pos":129,"Text":"Cg=="}]}} <nil>
