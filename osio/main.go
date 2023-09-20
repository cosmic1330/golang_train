package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 取得目前路徑
func GetPath() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current working directory:", cwd)
}

// 判斷文件是否存在
func FileExist(path string) bool {
	fileInfo, err := os.Stat(path) // os.Stat獲取文件信息
	if err != nil {
		return os.IsExist(err) // IsExist返回一個布爾值，指示錯誤是否表示文件或目錄已存在。
	}
	return !fileInfo.IsDir()
}

// 創建文件夾
func CreateDir(path string) bool {
	paths := strings.Split(path, "/")

	// 去除最後一個paths
	paths[len(paths)-1] = ""

	for i, v := range paths {
		if i == len(paths)-1 {
			break
		}
		if i != 0 {
			paths[len(paths)-1] += "/"
		}
		paths[len(paths)-1] += v
	}
	fmt.Println(paths)
	return os.MkdirAll((paths[len(paths)-1]), 0777) == nil
}

// 判斷文件夾是否存在
func DirExist(path string) bool {
	dirEntrys, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, v := range dirEntrys {
		if v.IsDir() {
			fmt.Println("inner:", v.Name()+"/")
			continue
		}
		fmt.Println("ineer:", v.Name())
	}
	return true
}

// 文件讀寫 (一行一行讀取)
func FileReadWrite() {
	f5, err := os.OpenFile("osio/f5.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f5.Close()
	writer := bufio.NewWriter(f5) // 創建寫入器
	fmt.Println("size=", writer.Size()) // 查看緩衝區容量
	writer.WriteString("hello world\n") // 寫入字符串
	fmt.Println("available", writer.Available()) // 查看寫入器的緩衝區未使用的字節數
	fmt.Println("buffered", writer.Buffered()) // 查看緩衝區中已使用的字節數
	writer.Reset(f5) // 重置寫入器
	for i := 1; i <= 2; i++ {
		fileName := fmt.Sprintf("f%v", i)
		data, err := os.ReadFile("osio/"+fileName)
		if err != nil {
			panic(err)
		}
		data = append(data, '\n')
		writer.Write(data) // bufw.Write([]byte)將數據寫入緩衝區

	}
	writer.Flush() // 將緩衝區的數據寫入硬碟中的文件
}

func main() {
	// 取得目前路徑
	GetPath()
	// 判斷文件是否存在
	fmt.Println(FileExist("test/test/test/hi.go"))
	// 創建文件夾
	CreateDir("test/test/test/hi.go")
	// 判斷文件夾是否存在
	fmt.Println(DirExist("test"))
	// 刪除文件夾
	os.RemoveAll("test")

	/* 文件操作 */
	file, err := os.OpenFile("osio/123.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	// 寫入
	file.WriteString("hello world from OpenFile\n")
	defer file.Close() // 關閉文件

	// 無緩衝讀寫 （適合小文件）
	data, err := os.ReadFile("osio/123.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// 有緩衝讀寫 （適合大文件）
	// io.EOF // End-of-file 表示文件結尾
	// io.Discard // 丟棄所有讀取的數據
	// bufio.NewReader(file) // 創建默認大小的讀取緩衝區
	// bufio.NewReaderSize(file, 1024) // 創建指定大小的讀取緩衝區


	// 合併文件
	FileReadWrite()
}
