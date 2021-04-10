package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"text/template"
	"text/template/parse"
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
	if err != nil || tmp == nil {
		panic(err)
	}
	d, e := json.Marshal(tmp)
	fmt.Println(string(d), e)
	walkTree(tmp.Tree)

}
func walkTree(t *parse.Tree) {
	if t != nil {
		treeSet := getUnExportedField(t, "treeSet")
		fmt.Printf("%#v", treeSet)

		/*if ts, ok := treeSet.Interface().(map[string]*parse.Tree); ok {
			for k, v := range ts {
				fmt.Println(k)
				walkTree(v)
			}
		}
		*/
		/*
			for i := 0; i < treeSet.NumField(); i++ {
				val := treeSet.Field(i)
				fmt.Println(val)
			}
		*/
		fmt.Println(reflect.TypeOf(treeSet).Key())
		fmt.Println(reflect.TypeOf(treeSet).Elem())

	}
}

func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}

//{"Name":"text","ParseName":"text","Root":{"NodeType":11,"Pos":0,"Nodes":[{"NodeType":0,"Pos":129,"Text":"Cg=="}]}} <nil>
