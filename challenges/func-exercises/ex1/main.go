package main

import (
	"fmt"
)

func foo() int {
	return 88
}

func bar() (x int, s string) {
	x = 55
	s = "returned string"
	return x, s
}

func main() {

	fmt.Println(foo())
	fmt.Println(bar())
}
