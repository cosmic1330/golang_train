package main

import "fmt"

func main() {
	/* 從陣列中引用 Slice */
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var slice1 []int = arr[1:4] // [2, 3, 4]
	slice1[0] = 10
	fmt.Println("arr:", arr)       // [1, 10, 3, 4, 5]
	fmt.Println("slice1:", slice1) // [10, 3, 4]
	fmt.Println()

	/* 創建Slice */
	// case1
	makeSlice := make([]int, 5, 10) // make([]Type, len, cap) 初始值為0
	makeSlice[0] = 1000
	fmt.Println("makeSlice:", makeSlice)
	// case2
	var x []int = []int{1, 2, 3, 4, 5} // 自動創建底層陣列
	fmt.Println("x:", x)
	// case3
	var slice2 []int // 創建一個空的 Slice
	fmt.Println("slice2:", slice2)
	fmt.Println()

	/* 底層創建新的陣列，不再引用原陣列 */
	slice2 = append(x, 1, 2, 3)
	fmt.Println("x:", x)
	fmt.Println("slice2:", slice2)
	fmt.Println()

	/* 加入其他切片 */
	slice3 := append(slice2, slice1...)
	fmt.Println("slice3:", slice3)
	fmt.Println()

	/* 複製切片 */
	slice4 := []int{1,1}
	copy(slice4,slice3)
	fmt.Println("slice4:", slice4) // slice4長度只有2，只會複製slice3的前兩個元素
	fmt.Println()

	/* 使用索引訪問 Slice 元素 */
	fmt.Println("First Element:", slice2[0])
	fmt.Println("Last Element:", x[len(x)-1])
	fmt.Println("Get length of Slice:", len(makeSlice))
	fmt.Println("Get capacity of Slice:", cap(makeSlice))
	fmt.Println("Get range:", slice2[1:])
	fmt.Println()

	/* 遍歷 Slice */
	// case1
	for i := 0; i < len(makeSlice); i++ {
		fmt.Printf("case1 Index: %d, Value: %d\n", i, makeSlice[i])
	}
	// case2
	for i, value := range makeSlice {
		fmt.Printf("case2 Index: %d, Value: %d\n", i, value)
	}
}
