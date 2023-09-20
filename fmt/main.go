package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func GetUser(name string, age int) string {
	return name + " " + strconv.Itoa(age)
}

func main() {
	/*
		Println() 會在最後面加上換行符號
		Printf() 打印出格式化的字符串
	*/
	user := GetUser("xiaoming", 13)
	// 顯示類型
	fmt.Printf("%T", user)
	fmt.Println()
	// 顯示預設值
	fmt.Printf("%v", user)
	fmt.Println()
	// 顯示U整數
	fmt.Printf("%U", 123)
	fmt.Println()
	// 顯示浮點數
	fmt.Printf("%f", 3.14)
	fmt.Println()

	/* 型別轉換 */
	i1 := 456
	s1 := "test"
	s2 := fmt.Sprintf("%d@%s", i1, s1) // 整數轉字串其他類型轉字串
	fmt.Println(s2)
	fmt.Printf("s2=%T", s2)

	var s3 string
	var i3 int
	n, err := fmt.Sscanf(s2, "%d@%s", &i3, &s3) // 字串轉其他類型
	fmt.Println(n)
	fmt.Println("i3=", i3)
	fmt.Println("s3=", s3)
	if err != nil {
		panic(err)
	}


	/* strconv 型別轉換 */
	var s4 = strconv.FormatInt(123, 4);
	println("s4", s4)
	var s5, err5 = strconv.Atoi("123") // 字串轉整數
	println("s5", s5, err5)
	var s6 = strconv.Itoa(123) // 整數轉字串
	println("s6", s6)
	var s7, err7 = strconv.ParseInt("456", 10, 64) // 字串轉整數
	println("s7", s7, err7)
	var s8, err8 = strconv.ParseUint("789", 10, 64) // 字串轉整數 64位元
	println("s8", s8, err8)
}
