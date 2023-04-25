package main

import "fmt"

func main() {

	for sum := 1; sum < 5000; {
		sum += sum
		for i := 0; i != sum; i++ {
			fmt.Printf("%d,", i)
		}

	}
	fmt.Printf("\n")
}
