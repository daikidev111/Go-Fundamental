package main

import (
	"fmt"
)

var pl = fmt.Println

// func useFunc(fun func(int, int) int, x, y, int) {
// 	pl(" Answer: ", (fun(x, y))) 
// }

func sumVals(x, y int) int {return x+y}


func main() {

	intSum := func(x, u int)int {return x + u}
	pl("5 + 4 = :", intSum(5, 4))

	sampl1 := 1
	changeVar := func() {
		sampl1 += 1
	}
	changeVar()
	pl("sampl1 = ", sampl1)

	// useFunc(sumVals(1, 2))
}