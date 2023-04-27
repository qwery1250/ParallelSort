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

	// 生成100個亂數
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = rand.Intn(1000)
	}

	// 排序
	startTime := time.Now()
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	endTime := time.Now()

	// 計算程序執行時間和最大記憶體佔用量
	duration := endTime.Sub(startTime)
	runtime.ReadMemStats(&mem)
	endMem := mem.Alloc
	maxMem := endMem - startMem

	fmt.Printf("排序後的數組: %v\n", nums)
	fmt.Printf("程式執行時間: %v\n", duration)
	fmt.Printf("最大記憶體佔用量: %d bytes\n", maxMem)
}
