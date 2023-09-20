package main

import (
	"bufio"
	"fmt"
	"os"
)

// 取得目前路徑
func GetPath() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current working directory:", cwd)
}

// 文件讀寫 (一行一行讀取)
func FileReadWrite() {
	f5, err := os.OpenFile("file/combin.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f5.Close()
	writer := bufio.NewWriter(f5)                // 創建寫入器
	fmt.Println("size=", writer.Size())          // 查看緩衝區容量
	writer.WriteString("hello world\n")          // 寫入字符串
	fmt.Println("available", writer.Available()) // 查看寫入器的緩衝區未使用的字節數
	fmt.Println("buffered", writer.Buffered())   // 查看緩衝區中已使用的字節數
	writer.Reset(f5)                             // 重置寫入器
	for i := 1; i <= 2; i++ {
		fileName := fmt.Sprintf("f%v", i)
		data, err := os.ReadFile("file/"+fileName)
		if err != nil {
			panic(err)
		}
		data = append(data, '\n')
		writer.Write(data) // bufw.Write([]byte)將數據寫入緩衝區

	}
	writer.Flush() // 將緩衝區的數據寫入硬碟中的文件
}

func main() {

	// 有緩衝讀寫 （適合大文件）
	// io.EOF // End-of-file 表示文件結尾
	// io.Discard // 丟棄所有讀取的數據
	// bufio.NewReader(file) // 創建默認大小的讀取緩衝區
	// bufio.NewReaderSize(file, 1024) // 創建指定大小的讀取緩衝區

	// 合併文件
	FileReadWrite()
}
