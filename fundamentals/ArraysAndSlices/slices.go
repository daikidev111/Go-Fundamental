package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// length and capacity may be different
	fmt.Println(len(a), cap(a))

	b := a[:]   // slice all
	c := a[3:]  // slice from the 4th to the end
	d := a[:6]  // slice first 6th
	e := a[3:6] // slice from the 4th to the 6th
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// you can use slicing for array too

	// MAKE
	// 1st param: type of object you want to create
	// 2nd param: length of the slice
	// 3rd param: capacity size to avoid expensive copy for creating the larger array
	make_arr := make([]int, 3, 100)
	fmt.Println(make_arr, len(make_arr), cap(make_arr)) // return [0, 0, 0]

	// slice concatenation with another slice
	make_arr = append(make_arr, []int{1, 2, 3, 4}...)
	fmt.Println(make_arr, len(make_arr), cap(make_arr)) // return [0, 0, 0]

}
