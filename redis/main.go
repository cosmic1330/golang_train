package main

import (
	"context"
	"fmt"
	"time"

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
	cli := client.Do(ctx, "SET", "foo1", "taibai")
	res1, err1 := cli.Result()
	if err1 != nil {
		if err1 == redis.Nil {
			fmt.Println("foo1 does not exist")
		} else {
			panic(err1)
		}
	} else {
		fmt.Println("foo1 save status", res1) // "OK
		res2, err2 := client.Do(ctx, "GET", "foo1").Result()
		fmt.Println("foo1", res2, err2)
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

	/* 類型處理 */
	client.Do(ctx, "SET", "bool1", true)
	client.Do(ctx, "SET", "bool2", 0)
	val3, err3 := client.MGet(ctx, "bool1", "bool2").Result()
	fmt.Println("bool", val3, err3)
	val4, err4 := client.Do(ctx, "MGET", "bool1", "bool2").BoolSlice()
	fmt.Println("bool slice", val4, err4)
	client.Do(ctx, "SET", "time1", time.Now())
	client.Do(ctx, "SET", "time2", time.Now().Unix())
	val5, err5 := client.Do(ctx, "MGET", "time1", "time2").Result()
	fmt.Println("time", val5, err5)

	/* pipeline 管道*/
	pipe := client.Pipeline()
	int1 := pipe.Get(ctx, "int1")
	fmt.Println("執行前int1", int1)
	for i := 0; i < 10; i++ {
		pipe.Set(ctx, "int1", i, 0)

	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("執行後int1", int1)

	/* transaction 事務*/
	client.Set(ctx, "user1", 0, 0)
	// 搶紅包
	handle := func(user string) {
		transactionErr := client.Watch(ctx, func(tx *redis.Tx) error {
			pipe := tx.Pipeline()
			err := pipe.IncrBy(ctx, user, 100).Err()
			if err != nil {
				return err
			}
			_, err = pipe.Exec(ctx)
			return err
		}, user)
		if transactionErr != nil {
			fmt.Println("搶紅包失敗", transactionErr)
		} else {
			fmt.Println("搶紅包成功", user)
		}
	}
	handle("user1")
	handle("user1")

	/* 遍歷 */
	iter := client.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Printf("key=%v, value=%v\n", iter.Val(), client.Get(ctx, iter.Val()).Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	/* hash */
	client.HSet(ctx, "hash1", "name", "johndoe")
	client.HSet(ctx, "hash1", "email", "johndoe@example.com")
	iter2 := client.HScan(ctx, "hash1", 0, "", 0).Iterator()
	for iter2.Next(ctx) {
		fmt.Printf("field=%v\n", iter2.Val())
	}
	/* 將Redis Hash 放到Go struct */
	type User struct {
		Username string `redis:"name"`
		Email    string `redis:"email"`
	}
	client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, "hash2", "name", "Jane")
		pipe.HSet(ctx, "hash2", "email", "jane@example.com")
		return nil
	})
	
	var user1 User
	if err := client.HGetAll(ctx, "hash1").Scan(&user1); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user1)
	
	var user2 User
	if err := client.HGetAll(ctx, "hash2").Scan(&user2); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user2)
	/* 關閉 */
	client.Close()
}
