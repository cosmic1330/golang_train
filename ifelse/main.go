package main

import (
	"fmt"
)

func main() {
	var age int = 18
	if age > 10 {
		fmt.Println("age > 10")
	} else if age > 20 {
		fmt.Println("age > 20")
	} else {
		fmt.Println("age < 10")
	}
}
