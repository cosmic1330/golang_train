package main

import (
	"fmt"
)

/*
switch 語法
switch 特性: 用于条件判断或各種類型及型別操作，接口 interface{} 型別判斷 variable.(type)，重點是會依照 case 順序依序執行。
*/
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

	var i interface{} ="foo"
	switch t := i.(type) {
    case int:
        println("i is interger", t)
    case string:
        println("i is string", t)
    case float64:
        println("i is float64", t)
    default:
        println("type not found")
    }
}
