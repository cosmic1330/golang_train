package main

import "fmt"


func main() {
	/* 自定義類型 */
	type myType map[string]int
	var a myType = myType{"a":1, "b":2}
	fmt.Printf("%T", a)
	fmt.Println()

	/* 類型別名 */
	type myInt = int
	var b myInt = 1
	fmt.Printf("%T", b)
}
