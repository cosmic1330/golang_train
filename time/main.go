package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now()) // 2021-08-15 15:50:00.000000 +0800 CST m=+0.000000001
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2021-08-15 15:50:00
	d1, err1 := time.ParseDuration("1h10m300s") // 從字串解析出時間
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(d1)
	time.Sleep(time.Second * 2) // 睡兩秒
	fmt.Println(time.Now().Add(time.Hour * 2)) // 現在時間加兩小時
	fmt.Println(time.Now().AddDate(1, 1, 1)) // 現在時間加一年一月一日
	fmt.Println(time.Now().AddDate(-1, -1, -1)) // 現在時間減一年一月一日
	var t = time.Now().Add(time.Hour * 2)
	fmt.Println(t.Sub(time.Now())) // 時間差
	fmt.Println(t.Date()) // 年月日
	fmt.Println(t.Year()) // 年
	fmt.Println(t.Month()) // 月
	fmt.Println(t.Day()) // 日
	fmt.Println(t.Hour()) // 時
	fmt.Println(t.Minute()) // 分
	fmt.Println(t.Second()) // 秒
	
	// 時間戳
	fmt.Println(time.Now().Unix()) // 1629022200
}
