package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter up to 10 integers separated by space:")

	// 读一整行
	var line string
	fmt.Scanln(&line)

	// 拆成字段
	fields := strings.Fields(line)

	// 转为 int slice，最多 10 个
	nums := make([]int, 0, 10)
	for i, s := range fields {
		if i >= 10 {
			break
		}
		n, err := strconv.Atoi(s)
		if err == nil {
			nums = append(nums, n)
		}
	}

	// 冒泡排序
	BubbleSort(nums)

	fmt.Println("Sorted:", nums)
}

// BubbleSort: 冒泡排序，原地修改 slice
func BubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

// Swap: 交换 arr[i] 和 arr[i+1]
func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}
