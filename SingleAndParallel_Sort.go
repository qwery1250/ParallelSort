package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

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

func parallelQuickSort(arr []int) []int {
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

	// 使用 WaitGroup 来等待并行任务完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 并行处理 less 和 more 数组
	go func() {
		defer wg.Done()
		runtime.Gosched()
		less = parallelQuickSort(less)
	}()

	go func() {
		defer wg.Done()
		runtime.Gosched()
		more = parallelQuickSort(more)
	}()

	// 等待并行任务完成
	wg.Wait()

	// 合并结果
	result := append(less, equal...)
	result = append(result, more...)
	return result
}

func main() {
	// 記錄程序開始前的記憶體佔用情況
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	startMem := mem.Alloc
	// 生成随机数数组
	arr := make([]int, 100000)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		arr[i] = rand.Intn(1000)
	}
	//fmt.Println("Original array: ", arr)
	// 使用并行快速排序
	startTime := time.Now()
	parallelQuickSort(arr)
	//fmt.Println("Sorted array (parallel): ", sorted)
	fmt.Println("Time taken (parallel): ", time.Since(startTime))
	endTime := time.Now()
	runtime.ReadMemStats(&mem)
	endMem := mem.Alloc
	maxMem := endMem - startMem

	// 計算程序執行時間和最大記憶體佔用量
	duration := endTime.Sub(startTime)
	fmt.Printf("程式執行時間: %v\n", duration)
	fmt.Printf("最大記憶體佔用量: %d bytes\n", maxMem)

	// 使用单线程快速排序
	startTime = time.Now()
	quickSort(arr)
	//fmt.Println("Sorted array (single thread): ", sorted)
	fmt.Println("Time taken (single thread): ", time.Since(startTime))
	endTime = time.Now()
	runtime.ReadMemStats(&mem)
	endMem = mem.Alloc
	maxMem = endMem - startMem

	// 計算程序執行時間和最大記憶體佔用量
	duration = endTime.Sub(startTime)
	fmt.Printf("程式執行時間: %v\n", duration)
	fmt.Printf("最大記憶體佔用量: %d bytes\n", maxMem)

}
