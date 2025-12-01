package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	// 提示用户输入文件名
	var filename string
	fmt.Print("input file name: ")
	fmt.Scan(&filename)

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file open failed:", err)
		return
	}
	defer file.Close()

	// 创建一个 Name 结构体的切片
	var names []Name

	// 获取文件信息以确定文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed:", err)
		return
	}

	// 读取文件内容到一个字节切片
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("read file failed:", err)
		return
	}

	// 将字节切片转换为字符串，并按行分割
	fileContent := string(buffer)
	lines := strings.Split(fileContent, "\n")

	// 处理每一行
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			fname := parts[0]
			lname := parts[1]
			// 将名字和姓氏存入结构体并添加到切片中
			names = append(names, Name{Fname: fname, Lname: lname})
		}
	}

	// 遍历结构体切片并打印每个结构体中的名字和姓氏
	for _, name := range names {
		fmt.Printf("fname: %s, lname: %s\n", name.Fname, name.Lname)
	}
}