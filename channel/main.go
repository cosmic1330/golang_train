package main

import (
    "fmt"
    "sync"
)

/*
   當goroutine使用相同變數進行讀寫時，會被覆蓋因此需要使用channel來避免這種情況。
   如果沒有兩個goroutine需要交換訊息，則使用lock和wait group會更好。
*/

func makeHTTPRequests(requestChan <-chan string) {
	for url := range requestChan {
		fmt.Println("request url: ", url)
	}
}

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sending: %d\n", i)
		ch <- i // 将数据发送到通道
	}
	close(ch) // 关闭通道，通知接收方不会再有数据
}

func receiver(ch chan int) {
	for {
		val, ok := <-ch // 从通道接收数据
		if !ok {
			fmt.Println("Channel is closed. Receiver exiting.")
			return
		}
		fmt.Printf("Received: %d\n", val)
	}
}

func main() {
	/* 有緩衝區的channel */
	var wg sync.WaitGroup
	maxConcurrentRequests := 10
	chan1 := make(chan string, maxConcurrentRequests)
	for i := 0; i < maxConcurrentRequests; i++ {
		wg.Add(1)
		go makeHTTPRequests(chan1)
	}

	// Add tasks to the channel
	for i := 0; i < 50; i++ {
		chan1 <- "http://example.com?i=" + fmt.Sprintf("%d", i)
	}
	close(chan1)
	// Wait for all goroutines to finish
    wg.Wait()


	/* 無緩衝區的channel */
	msgCh := make(chan int) // 创建一个无缓冲通道
	go sender(msgCh)
	go receiver(msgCh)


	/* 如果你有多個 channel 需要讀取，而讀取是不間斷的，就必須使用 for + select 機制來實現 */
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		fruits := []string{"apple", "banana", "cherry", "date", "fig"}
		for _, fruit := range fruits {
			ch2 <- fruit
		}
		close(ch2)
	}()

	// 使用 for + select 从多个通道读取数据
	for {
		select {
		case num, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 is closed. Exiting.")
				return
			}
			fmt.Printf("Received from ch1: %d\n", num)
		case fruit, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 is closed. Exiting.")
				return
			}
			fmt.Printf("Received from ch2: %s\n", fruit)
		}
	}
}
