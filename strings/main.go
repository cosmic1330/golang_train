package main

import (
	"fmt"
	"strings"
)

func main() {
	bool1 :=strings.Contains("test", "es") // 字串是否包含
	fmt.Println("bool1", bool1)
	strings.Count("test", "t") // 字串出現次數
	strings.Index("test", "t") // 字串第一次出現位置
	strings.LastIndex("test", "t") // 字串最後一次出現位置
	strings.Replace("test", "t", "a", 2) // 字串替換前兩個 <-1替換全部>
	strings.ReplaceAll("test", "t", "a") // 字串替換全部
	strings.Split("test", "e") // 字串分割
	strings.ToLower("TEST") // 字串轉小寫
	strings.ToUpper("test") // 字串轉大寫
	strings.Repeat("test", 2) // 字串重複兩次
	strings.HasPrefix("test", "te") // 是否存在字串前綴
	strings.Fields("test test") // 字串分割成切片
	strings.Join([]string{"test", "test"}, " ") // 切片合併成字串
	strings.Trim("test", "t") // 字串去除前後指定字元
	strings.TrimSpace(" test ") // 字串去除前後空白
}
