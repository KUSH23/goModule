package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}

}

func main() {
	var s1 Speaker
	// var d1 Dog = Dog{"Brian"}
	var d1 *Dog
	s1 = d1
	s1.Speak()
}
