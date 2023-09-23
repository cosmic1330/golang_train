package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	/* 
		Context 
		作用: 允許在程序中傳遞一個request相關的值，並設置一個timeout，當超過這個timeout時，自動取消這個request
		使用場景:
			1. 用於控制goroutine的生命週期
			2. 中斷一個正在執行的goroutine後
	*/
	ctx := context.Background() // 返回一个空的Context，永遠不會被取消，沒有有效期限

	// 方法一
	cli:=client.Do(ctx, "SET", "foo1", "taibai")
	res, err1 := cli.Result()
	if err1 != nil {
		if err1 == redis.Nil {
			fmt.Println("foo1 does not exist")
		} else {
			panic(err1)
		}
	} else {
		fmt.Println("foo1", res)
	}
	
	// 方法二
	err := client.Set(ctx, "foo2", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo2", val)
}
