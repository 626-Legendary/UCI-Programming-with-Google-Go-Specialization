package main

import (
	"fmt"
	"strconv"
	"strings"
)

func bubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	// 初始 slice，容量 3，长度 0
	sli := make([]int, 0, 3)

	for {
		fmt.Print("Enter an integer (or X to exit): ")

		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)

		// 判断退出
		if strings.EqualFold(input, "x") {
			fmt.Println("Exiting.")
			break
		}

		// 使用 strconv 转整数
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid integer.")
			continue
		}

		sli = append(sli, num)

		bubbleSort(sli)

		fmt.Println("Sorted slice:", sli)
	}
}

/*
“为什么 slice 需要限制容量？容量的意义是什么？”
1. 容量（capacity）是 Go 的底层优化机制
2. 这是 Go 语言最重要的性能优化机制之一
3. 初始容量越小 → 越容易触发“扩容机制”→ 理解 slice 工作原理
*/