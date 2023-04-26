package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 設置亂數種子
	arrLen := 10
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(100) // 產生 0-99 的亂數整數
	}
	for _, n := range arr {
		fmt.Printf("%2d,", n)
	}
	fmt.Printf("array len %d\n", len(arr))
	fmt.Printf("max value: %d\n", max(arr))
}
