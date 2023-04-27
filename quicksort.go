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
	quickSort(nums)

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

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 选择中间值作为基准
	pivot := arr[len(arr)/2]

	// 分割成三个数组
	less := make([]int, 0)
	equal := make([]int, 0)
	more := make([]int, 0)
	for _, val := range arr {
		switch {
		case val < pivot:
			less = append(less, val)
		case val == pivot:
			equal = append(equal, val)
		case val > pivot:
			more = append(more, val)
		}
	}

	// 递归排序
	less = quickSort(less)
	more = quickSort(more)

	// 合并结果
	result := append(less, equal...)
	result = append(result, more...)
	return result
}
