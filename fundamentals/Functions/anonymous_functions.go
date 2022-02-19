package main

import "fmt"

func main() {

	// Basic code for anonymous function
	func() {
		fmt.Println("Hello Go!")
	}() // () to invoke the function immediately

	// what if you want to pass a value from the outer scope?
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i) // pass i
	}

	// you can also assign an anonymous function to a variable
	f := func() {
		fmt.Println("Hello World")
	}
	f()

	// another way
	var d func() = func() {
		fmt.Println("Hello World")
	}
	d()
}
