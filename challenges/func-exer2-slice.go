package main

import (
	"fmt"
)

func funcvar(x ...int) {
	fmt.Println(x)
	sum := 0
	for i, v := range x {

		sum += v
		fmt.Println(i, v, sum)
	}
}
func main() {

	funcvar(1, 50, 100, 200)
}
