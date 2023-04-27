package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 產生亂數數組
	arr := generateRandomArray(100000)
	fmt.Println("Original array: ", arr[:10], "...")

	// 單線程排序
	start := time.Now()
	bubbleSort(arr)
	duration := time.Since(start)
	fmt.Println("Sorted array (single thread): ", arr[:10], "... (total length: ", len(arr), ")")
	fmt.Println("Time taken (single thread): ", duration)

	// 使用平行運算排序
	start = time.Now()
	parallelBubbleSort(arr)
	duration = time.Since(start)
	fmt.Println("Sorted array (parallel): ", arr[:10], "... (total length: ", len(arr), ")")
	fmt.Println("Time taken (parallel): ", duration)
}

// 產生亂數數組
func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

// 單線程氣泡排序
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 使用平行運算排序
func parallelBubbleSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	done := make(chan bool)

	// 平分陣列，分別進行排序
	go func() {
		bubbleSort(arr[:n/2])
		done <- true
	}()
	go func() {
		bubbleSort(arr[n/2:])
		done <- true
	}()

	// 等待兩個排序完成
	<-done
	<-done

	// 合併已排序的兩個陣列
	merge(arr, n/2)

}

// 合併已排序的兩個陣列
func merge(arr []int, mid int) {
	n := len(arr)
	temp := make([]int, n)
	copy(temp, arr)

	/*宣告一個新的陣列，其長度為兩個待合併陣列的總和。
	分別從兩個待合併陣列中選出第一個元素，比較其大小，將較小的元素放入新的陣列中。
	持續比較兩個待合併陣列中的元素，直到其中一個陣列中的所有元素都已被放入新陣列中。
	將未放入新陣列中的另一個陣列中的元素依序放入新陣列中。*/
	i := 0
	j := mid
	for k := 0; k < n; k++ {
		if i >= mid {
			arr[k] = temp[j]
			j++
		} else if j >= n {
			arr[k] = temp[i]
			i++
		} else if temp[j] < temp[i] {
			arr[k] = temp[j]
			j++
		} else {
			arr[k] = temp[i]
			i++
		}
	}
}
