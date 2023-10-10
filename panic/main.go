package main

import (
	"fmt"
)
/*
	使用 panic 會對程序的效能產生一定的影響。當程序發生 panic 時，
	它會從当前的 goroutine 中返回并遍历调用栈，直到找到一个能够处理该 panic 的 recover 函数。
	不想使用 panic，你可以使用 error 类型来处理错误。
	在使用 goroutine 时，你可以使用 panic 和 recover 来处理错误，它不会影响其他 goroutine 的执行。
*/
// 程式出現錯誤或出現panic時，會中指程式的執行，並且會印出錯誤訊息
func main() {
	fmt.Println("start")
	// panic("panic message")
	sendPanic()
	fmt.Println("end")	
}

func sendPanic(){
	defer handlePanic()
	panic("panic message")
}

func handlePanic(){
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}