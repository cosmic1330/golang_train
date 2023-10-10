package main

import (
	"errors"
	"fmt"
)

// 自定義錯誤結構體
type DivisionError struct {
	message string
}

// ERROR接口的Error函式
func (z DivisionError) Error() string {
	return "錯誤:"
}
func main() {
	var number, divisor int
	// fmt.Scan(&number)
	// fmt.Scan(&divisor)
	number = 1 
	divisor = 0
	res, err := divide(number, divisor)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(res)
}

func divide(number, divisor int) (float32, error) {
	if divisor == 0 {
		err := errors.New("分母不能為0!")
		err = fmt.Errorf("分母不能為%d!", divisor)
		err = DivisionError{message: "分母不能為0!"}
		return 0.0, err
	}
	return float32(number / divisor), nil
}
