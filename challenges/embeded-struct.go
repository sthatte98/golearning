package main

import "fmt"

type vehicle struct {
	doors int
	color string
	//faviceflvr []string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

//var pslice []person

func main() {

	T1 := truck{

		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	S1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: false,
	}

	fmt.Println(T1.doors)
	fmt.Println(S1)

	/*
		// pslice = append(pslice, p0, p1)

		// for _, stru := range pslice {
		// 	fmt.Println(stru)
		// 	for _, flvr := range stru.faviceflvr {

		// 		fmt.Println(flvr)
		// 	}

		// }
	*/
}
