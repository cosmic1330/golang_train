package main

import (
	"case/pass"
	"fmt"
)

func main() {
	/* case1 base*/
	var x int = 3
	var xPtr *int = &x // 取得 x 的記憶體位址
	fmt.Println(xPtr)
	fmt.Println(*xPtr)

	/* case2 pass*/
	var a int = 10
	var aPtr *int = &a // 取得 x 的記憶體位址
	pass.Add(aPtr)
	fmt.Println(a)
}
