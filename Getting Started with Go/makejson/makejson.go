package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Print("Enter your first name: ")
	fmt.Scanln(&name) // 注意：Scanln 遇到空格会停，如果地址里有空格就需要更复杂的处理

	fmt.Print("Enter your address: ")
	fmt.Scanln(&address)

	// 使用 map 存储 name 和 address
	personMap := make(map[string]string)
	personMap["name"] = name
	personMap["address"] = address

	// 将 map 转为 JSON
	jsonBytes, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	// 打印 JSON 字符串
	fmt.Println(string(jsonBytes))
}
