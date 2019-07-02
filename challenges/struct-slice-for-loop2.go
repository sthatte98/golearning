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

	m := make(map[string]person)

	m[p0.last] = p0

	pslice = append(pslice, p0, p1)

	// populate the map with person's last name as an index and the struck for each p0 and p1 as values
	for _, stru := range pslice {
		m[stru.last] = stru
	}

	for lastname, mapStruct := range m {
		fmt.Println(lastname)
		for _, v := range mapStruct.faviceflvr {
			fmt.Println(v)
		}
	}
}
