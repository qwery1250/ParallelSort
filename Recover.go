package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USESR")

func main() {
	err := throwsPanic()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
func throwsPanic() (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error: %v", x)
		}
	}()
	if user == "" {
		panic("no value for $USER")
	}
	return
}
