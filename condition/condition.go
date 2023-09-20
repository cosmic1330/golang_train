package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const identifier int = 100
	fmt.Println("case1:", identifier)
	// for
	/* 2-1
		+=語句會分配一個新的字符串，併將老字符串連接起來的值賦予給它。
		而目標字符串的老字面值在得到新值以後就失去了用處，這些臨時值會被Go語言的垃圾收集器幹掉。
		如果不斷連接的數據量很大，那麽下面這種操作就是成本非常高的操作。 
	*/
	var s, sep string = "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println("case2:", s)
	/* 
		2-2 推薦做法是使用strings包中的Join函數來完成相同的功能，它的性能要比前面的寫法高很多。
	*/
	fmt.Println("case2:", strings.Join(os.Args[1:], " "))

	// as while
	num := 1
	for num < 10 {
		fmt.Println("case3:", num)
		num++
	}
}
