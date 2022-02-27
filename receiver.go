package main

import (
	"coursera/data"
	"fmt"
)

type MyInt int

func (x MyInt) Double() int {
	return int(x * 2)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())
	fmt.Println(data.X)
}
