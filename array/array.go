package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)
	// 賦值
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2, len(arr2))
	// 自動判斷長度
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3, len(arr3))
	// 遍歷
	var arr4 [5]int = [5]int{5, 4, 3, 2, 1}
	for i := 0; i < len(arr4); i++ {
		fmt.Print(arr4[i])
	}
	fmt.Println()
	// 多維陣列
	var arr5 [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6}}
	fmt.Println(arr5)
	// 多維陣列遍歷
	for i := 0; i < len(arr5); i++ {
		for j := 0; j < len(arr5[i]); j++ {
			fmt.Print(arr5[i][j], " ")
		}
		fmt.Println()
	}
	
}
