package main

import "fmt"

func main() {
	/* 創建map */
	map1:= map[string]string{"a":"1","b":"2"}
	fmt.Println(map1)
	map2 := make(map[string]string, 2)
	map2["a"] = "1"
	map2["b"] = "2"
	map2["c"] = "3"
	fmt.Println(map2)
	fmt.Println()

	/* 查找map */
	value, ok := map2["a"]
	if ok {
		fmt.Println("value:", value)
	} else {
		fmt.Println("key not exist")
	}
	fmt.Println()

	/* 刪除map */
	delete(map2, "a")
	fmt.Println(map2)

	/* 遍歷map */
	for key, value := range map2 {
		fmt.Println("key:", key, "value:", value)
	}
	fmt.Println()
}
