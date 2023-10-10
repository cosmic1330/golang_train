package main

import (
	"fmt"
)

func main(){
	// 後進先出
	defer fmt.Println("go") // 延遲執行
	defer fmt.Println("python") // 延遲執行
}