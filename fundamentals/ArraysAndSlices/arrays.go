package main

import (
	"fmt"
)

func main() {
	// declare an array with a fixed number of elements
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v \n", grades)

	// delclare an array with implicit number of elements
	grades_flex := [...]int{97, 85, 93, 94}
	fmt.Printf("Grades: %v \n", grades_flex)

	// array that contains 0 by not delclaring the elements
	var students [3]int
	fmt.Printf("student: %v \n", students)

	// element substitution
	students[0] = 100
	fmt.Printf("first student: %v \n", students[0])

	//creating a matrix with 0s
	var matrix [3][3]int
	fmt.Printf("matrix: %v \n", matrix)

	//assign array to a different var
	arr := [...]int{1, 2, 3}
	// this will copy the entire array and slow down the behaviour
	// arr_another := arr
	// How to avoid it
	arr_another := &arr // point to the arr data
	arr_another[0] = 100
	fmt.Println("swap result", *arr_another)
	// arr elem will change too if you use pointer assignment
	fmt.Println("swap result", arr)

	arr_cap_change := []int{}
	fmt.Println(arr_cap_change)
	fmt.Printf("intial length %v\n", len(arr_cap_change))
	fmt.Printf("initial cap %v\n", cap(arr_cap_change))

	arr_cap_change = append(arr_cap_change, 1)
	// after append operation, it cannot fit into a zero element array, so
	// it assigns an array for us; copies an array and make a larger size
	// this could be expensive for large array, leading to a lot of memory consumption
	fmt.Println(arr_cap_change)
	fmt.Printf("Now length %v\n", len(arr_cap_change))
	fmt.Printf("Now cap %v\n", cap(arr_cap_change))

}
