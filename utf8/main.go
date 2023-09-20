package main

import (
	"unicode/utf8"
)

func main() {
	utf8.RuneCountInString("test") // 字串長度
	utf8.Valid([]byte("test")) // 字串是否合法
	utf8.ValidRune('t') // 字元是否合法
	utf8.DecodeRune([]byte("test")) // 字串轉字元
	utf8.DecodeLastRune([]byte("test")) // 字串轉字元
}
