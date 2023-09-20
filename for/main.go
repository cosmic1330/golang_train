package main

import (
	"fmt"
)

func main() {
	// 無限循環
	i := 1
	for {
		fmt.Print(i)
		i++
		if i > 10 {
			fmt.Println()
			break
		}
	}
	// 條件循環
	i = 1
	for i < 11 {
		fmt.Print(i)
		i++
	}
	fmt.Println()
	// for
	for i := 1; i < 11; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	// for in for
	for i := 1; i <= 3; i++ {
		for j := 1; j < 11; j++ {
			fmt.Print(j, "\t")
			if j == 2 {
				break // 只會中指內部循環
			}
		}
	}
	fmt.Println()
	// label
label:
	for i := 1; i <= 3; i++ {
		for j := 1; j < 11; j++ {
			fmt.Print(j, "\t")
			if j == 2 {
				break label // 中斷指定循環
			}
		}
	}
	fmt.Println()
	// goto
	fmt.Println("goto")
	i = 1
	for i < 11 {
		fmt.Print("+")
		i++
		if i == 3 {
			goto label2
		}
	}
label2:
	fmt.Println("label1")
	fmt.Println()
}
