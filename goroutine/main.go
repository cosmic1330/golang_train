package main

import (
	"fmt"
	"sync"
)

var (
	count int
	lock sync.Mutex
)
func findPrime(n int){
	// 找出 1~n 中的質數
	for i := 2; i < n; i++ {
		if n % i == 0 {
			// 不是質數
			return
		}
	}
	// 是質數
	// fmt.Println("prime", n)
	count++
}

// 使用協程會影響共同變數count的正確性，必須使用sync.Mutex解決
func findPrimeUseLock(n int){
	// 找出 1~n 中的質數
	for i := 2; i < n; i++ {
		if n % i == 0 {
			// 不是質數
			return
		}
	}
	lock.Lock()
	count++
	lock.Unlock()
}

func main() {
	// case 1
	for i := 2; i <= 100000; i++ {
		go findPrime(i)
	}
	fmt.Println(count) // count 不正確，應該要是 9592

	// case 2 use lock （但use lock的時間和不使用協程差不多）
	// for i := 2; i <= 100000; i++ {
	// 	go findPrimeUseLock(i)
	// }
	// fmt.Println(count)
}
