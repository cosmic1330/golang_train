package main

import "fmt"

func pushNum(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func main() {
	/* 有緩衝區的channel */
	var c1 = make(chan int, 100)
	// 存入
	c1 <- 1
	c1 <- 2
	// 取出
	v, ok := <-c1
	fmt.Println("c1", v, ok)
	fmt.Println("c1", <-c1)

	// 關閉
	close(c1)

	/* 無緩衝區的channel */
	// 取出 for不需要關閉管道
	var c2 = make(chan int)
	go pushNum(c2)
	fmt.Println()
	for v := range c2 {
		fmt.Printf("c2: %v\t",v)
	}

	// 取出2 
	fmt.Println()
	fmt.Println()
	var c3 = make(chan int)
	go pushNum(c3)
	for {
		v, ok := <-c3
		if !ok {
			break
		}
		fmt.Printf("c3: %v<%v>\t", v, ok)
	}
}
