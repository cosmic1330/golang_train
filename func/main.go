package main

import (
	"fmt"
)

func getRes(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
}

// 匿名函數
var getRes2 = func(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
}

// defer
func deferUtil(other ...int) func(int) int {
	i := 0
	return func(a int) int {
		i++
		fmt.Println("a=", a)
		fmt.Println("i=", i)
		return i
	}
}

func runDefer() {
	f := deferUtil()
	defer f(100)
	defer f(f(200))
	defer f(300)
	/*
		執行順序
		f(200)
		f(300)
		f(res) // res為f(200)回傳值
		f(100)
	*/
}

// defer cover 錯誤捕捉
func DeferRecover() int{
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover:", err)
		}
	}()
	n := 0
	res := 3 / n
	return res
}

func main() {
	sum, diff := getRes(1, 2)
	sum2, diff2 := getRes2(2, 2)
	fmt.Println("sum", sum, "diff", diff)
	fmt.Println("sum2", sum2, "diff2", diff2)
	runDefer()
	DeferRecover()
}
