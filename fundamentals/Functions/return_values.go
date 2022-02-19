package main

import "fmt"

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", s)

	s_pointer := sumPointer(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s_pointer)

	s_named := sumNamed(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", s_named)

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

// How to return a value delcare what to return next to param
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// return local variable as a pointer
func sumPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	// you do not need to worry about freeing the memory and mess up the return value
	// As this is handled by the heap memory
	return &result
}

// How to return a named variable
func sumNamed(values ...int) (result_named int) {
	fmt.Println(values)
	for _, v := range values {
		result_named += v
	}

	// we do not have to specify the return variable as it is implicityly returned
	return
}

// How to return multiple values
func divide(a, b float64) (float64, error) {
	// we should not use panic as a general course of action
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero") // pass error object
	}
	return a / b, nil // pass nil
}
