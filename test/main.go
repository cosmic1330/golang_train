package main

import (
	"fmt"
	"test/method"
)

// https://blog.wu-boy.com/2018/06/how-to-write-benchmark-in-go/

func main() {
	fmt.Println(method.IsNegative(-1))
}
