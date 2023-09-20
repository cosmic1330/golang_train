package main

import (
	"utils/mytool"
)

func init() {
	println("main init")
}

func init() {
	println("main2 init")
}

func main() {
	x := 3
	mytool.Add(&x)
	println(x)
}
