package main

import (
	"fmt"
)

func main() {
	var age int = 19
	switch {
		case age == 18:
			fmt.Println("age = 18")
		case age < 18:
			fmt.Println("age < 18")
		case age > 18:
			fmt.Println("age > 18")
		default:
			fmt.Println("格式錯誤")
	}

}
