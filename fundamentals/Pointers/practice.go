package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal(ptr *int) {
	*ptr+=1
}

func dblArrVals(arr *[4]int) {
	for x:=0; x < 4; x++ {
		arr[x] *= 2
	}
}

func main() {
	num := 5	
	pl("num before func", num)
	changeVal(&num)
	pl("num after func", num)

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl("array", pArr)
}