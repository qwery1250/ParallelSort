package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// 記錄程序開始前的記憶體佔用情況
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	startMem := mem.Alloc

	// 產生 100 個亂數
	nums := make([]int, 100000)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		nums[i] = rand.Intn(1000)
	}

	// 使用快速排序排序
	startTime := time.Now()
	quicksort(nums)

	endTime := time.Now()

	// 計算程序執行時間和最大記憶體佔用量
	duration := endTime.Sub(startTime)
	runtime.ReadMemStats(&mem)
	endMem := mem.Alloc
	maxMem := endMem - startMem

	//fmt.Printf("排序後的數組: %v\n", nums)
	fmt.Printf("程式執行時間: %v\n", duration)
	fmt.Printf("最大記憶體佔用量: %d bytes\n", maxMem)
}

func quicksort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	pivot := nums[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for _, n := range nums[1:] {
		if n <= pivot {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}

	quicksort(left)
	quicksort(right)

	copy(nums, append(append(left, pivot), right...))
}
