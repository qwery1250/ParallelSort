package main

import (
	"fmt"
)

type person struct {
	height float32
	weight float32
}

func (p person) GetBMI() float32 {

	bmi := p.weight / p.height / p.height
	return bmi
}
func (p *person) setWeight(w float32) {

	p.weight = w
}
func main() {
	p := person{1.68, 73}
	p.setWeight(100)

	fmt.Println(p.GetBMI())
}
