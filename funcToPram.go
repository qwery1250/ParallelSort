package main

import "fmt"

func filter(numbers []int, f func(int) bool) []int {
	var result []int
	for _, n := range numbers {
		if f(n) {
			result = append(result, n)
		}
	}
	return result
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 == 1
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(numbers, isEven)
	odds := filter(numbers, isOdd)
	fmt.Println("Even numbers:", evens)
	fmt.Println("Odd numbers:", odds)
}
