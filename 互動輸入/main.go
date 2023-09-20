package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	/* 接收參數 */
	// go run main.go 1 2 3
	fmt.Println("接收到了", len(os.Args), "個參數")
	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	// flag
	//  go run main.go -mode -u="mary" -f="required" "not flag"
	userMode := flag.Bool("mode", false, "使用者模式")
	var userName string
	flag.StringVar(&userName, "u", "", "使用者名稱")
	flag.Func("f", "必須輸入", func(s string) error {
		fmt.Println("flagFunc", s)
		return nil
	})
	flag.Parse()
	fmt.Println("userMode", *userMode)
	fmt.Println("userName", userName)
	// 顯示flag後沒有帶flag的參數
	for i, v := range flag.Args() {
		fmt.Println("flag", i, v)
	}

	/* 事後輸入參數 */
	var a int
	var b int
	// fmt.Print("Enter a number:")
	// fmt.Scanln(&a)
	// fmt.Print("Enter b number:")
	// fmt.Scanln(&b)
	fmt.Print("Enter a & b number: 用空格隔開")
	fmt.Scanln(&a, &b)
	var result int = a + b
	fmt.Println(result)
}
