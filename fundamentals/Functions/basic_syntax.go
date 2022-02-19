package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello", i)
	}
	greeting := "Hello"
	name := "World"
	sayGreeting(greeting, name)
	fmt.Println(name) // 1. Still Remain As "World"
	sayGreetingPointer(&greeting, &name)
	fmt.Println(name) // 2. Changed it to Ted!
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println(idx)
}

// greeting and name are the same data type, so you can ignore type declaration for greeting
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)

	// what happens to the main func name variable after assigning new value? -> 1
	name = "Ted"
	fmt.Println(greeting, name)

}

// what is the difference between passing pointers and variables?
// -> passing the variable is expensive as it copies every single time
// -> use pointer for data consistency and avoiding expensive copy
func sayGreetingPointer(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted" // what happens to the main func name variable after assigning new value? -> 2
	fmt.Println(*greeting, *name)
}
