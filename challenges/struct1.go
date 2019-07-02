package main

import "fmt"

type person struct {
	first      string
	last       string
	faviceflvr []string
}

var pslice []person

func main() {

	p0 := person{"james", "bond", []string{"rasberry", "vanila"}}
	p1 := person{"money", "penny", []string{"choco", "strawberry"}}

	pslice = append(pslice, p0, p1)

	for _, stru := range pslice {
		fmt.Println(stru)
		for _, flvr := range stru.faviceflvr {

			fmt.Println(flvr)
		}

	}

}
