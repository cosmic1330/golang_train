package main

import (
	"errors"
	"fmt"
)

func main() {
	// 自定義錯誤
	err1 := errors.New("自定義錯誤")
	fmt.Println(err1)

	err2 := fmt.Errorf("錯誤訊息: %s", "自定義錯誤")
	fmt.Println(err2)

	// 捕捉錯誤
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕捉到錯誤訊息:", err)
			fmt.Println("執行恢復後的程式")
		}
		// 这里是需要在recover后继续执行的代码
		fmt.Println("panic後的程式2")
	}()
	panic("錯誤訊息")
	
	fmt.Println("panic後的程式1") // 不會執行

}
