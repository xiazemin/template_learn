package main

import (
	"bytes"
	"fmt"
)

func getKey(GUID string) string {
	var buf = bytes.Buffer{}
	_, _ = buf.WriteString("prefix")
	_, _ = buf.WriteString(":")
	_, _ = buf.WriteString(GUID)
	return buf.String()
}

func main() {
	fmt.Println(getKey("xzmYUDNN"))
}
