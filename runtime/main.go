package main

import (
	"fmt"
	"runtime"
)


func main() {
	// 获取当前系统的架构、操作系统、CPU核数
	arch := runtime.GOARCH
	os := runtime.GOOS
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(4)
	fmt.Println(arch, os, cpu)

	// 線程控制
	routine := runtime.NumGoroutine()
	fmt.Println(routine)
}
