package main

import (
	"fmt"
	"sort"
)

type SortAny interface {
	SortSlice()
}

//my comment 1 2 3 4 5 6
type people []string
type nums []int

var studygroup = people{"Zeno", "John", "Al", "Jenny"}
var mynums = nums{7, 4, 8, 2, 9, 19, 12, 32, 3}

func (p people) SortSlice() {
	sort.Sort(sort.Reverse(sort.StringSlice(studygroup)))
	//sort.Strings(studygroup)
	fmt.Println(studygroup)
}
func (n nums) SortSlice() {
	sort.Ints(mynums)
	fmt.Println(mynums)
}

func SortThis(s SortAny) {
	s.SortSlice()
}
func main() {
	SortThis(studygroup)
	studygroup.SortSlice()
	SortThis(mynums)
	mynums.SortSlice()

}
