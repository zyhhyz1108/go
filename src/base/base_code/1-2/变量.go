package main

import "fmt"

func main() {
	var s1 string = "Hello"
	var s2 string
	var s3 = "Hello"
	s4 := "Hello"
	b := make([]int, 0, len(s2))
	map1 := make(map[string]int)
	map1["go"] = 99
	map1["python"] = 88
	map2 := map[string]int{
		"go":     99,
		"python": 88}
	fmt.Println(s1, s2, s3, s4, b, map1, map2)
}
