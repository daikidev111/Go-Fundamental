package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5)
}

// How to use variadic parameter
// But it has to be one and at the end in the param
func sum(values ...int) {
	fmt.Println(values) // shows that this parameter thing acts as a slice
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is: ", result)
}
