package main

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

func main() {
	msg := Msg{Info: &Info{ID: 123, Name: "test", Text: "msg context"}}
	data := getData(&msg)
	/*
		&main.Info{ID:123, Name:"test", Text:"msg context"}
		&reflect.rtype{size:0x8, ptrdata:0x8, hash:0x7d1166de, tflag:0x8, align:0x8, fieldAlign:0x8, kind:0x36, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x102c63450), gcdata:(*uint8)(0x102d80f30), str:12610, ptrToThis:0}
	*/
	typ := getType(msg)
	fmt.Printf("%#v \n%#v\n", data, typ)
	execute(t2, data, nil)
	//name:
	execute(t2, nil, data)
	//name:test
	execute(t, msg.Info, nil)
	//name:test
	execute(t, data, nil)
	//panic :executing "wrapper" at <Content>: can't evaluate field Name in type interface {}
	//fmt.Println(data.Name)
	execute(t, data, msg.Info)
	execute(t, data, data)
}

func execute(tpl string, data, dataIn interface{}) {

	bodyTemplate, err := template.New("bodyTemplate").Funcs(template.FuncMap{
		"Content": func() interface{} { return data },
	}).Parse(tpl)
	if err != nil {
		panic(err)
	}
	builder := &strings.Builder{}
	if err := bodyTemplate.ExecuteTemplate(builder, "wrapper", dataIn); err != nil {
		panic(err)
	}
	fmt.Println(builder.String())
}

type Msg struct {
	Info *Info
}

type Info struct {
	ID   int
	Name string
	Text string
}

func getData(msg *Msg) interface{} {
	return reflect.ValueOf(msg).Elem().FieldByName("Info")
}

func getType(msg Msg) interface{} {
	field, ok := reflect.TypeOf(msg).FieldByName("Info")
	if !ok {
		return nil
	}
	return field.Type
}

var t string = `{{ define "wrapper" }}
name:{{ Content.Name }}
end
{{ end }}
`

var t2 string = `{{ define "wrapper" }}
{{.}}
name:{{ .Name }}
end
{{ end }}
`
