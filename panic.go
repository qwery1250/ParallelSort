package main

import "fmt"

func divide(a, b int) (result int, err error) {
	defer func(msg string) {
		if p := recover(); p != nil {
			err = fmt.Errorf("%sdivision error: %v", msg, p)
		}
	}("這裡面是傳給defer的參數。\n")
	if b == 0 {
		panic("division by zero")
	}
	result = a / b
	return
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
