package main

import "fmt"

// 遞迴
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// 閉包
func Closure() func(int) int {
	i := 0
	return func(n int) int {
		i++
		fmt.Println("接收參數n=", n)
		fmt.Println("閉包執行次數i=", i)
		return i
	}
}

func main() {
	n := 5
	println(fibonacci(n))

	// 閉包
	f := Closure()
	f(5)
	f(4)
	f(3)
	f = Closure()
	f(2)

}
