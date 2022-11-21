package main

import (
	"fmt"
)

var pl = fmt.Println

type Tsp float64
type TBs float64
type ML float64


// associate method
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsp = %.2f ML \n", ml1)
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f ML \n", ml2)

	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f ml \n", tsp1, tsp1.ToMLs())
}