package main

import "fmt"

func getPath() (r []string) {
	for i := 0; i < 3; i++ {
		r = append(r, fmt.Sprintf("%d", i))
	}
	return r
}

func main() {
	intToPath := make(map[int][]string)
	var GUIDs []string
	for i := 1; i < 4; i++ {
		p := getPath()
		intToPath[i] = p
		fmt.Println(p)
		GUIDs = append(GUIDs, p...)
	}
	fmt.Println(intToPath)
	fmt.Println(GUIDs)
}
