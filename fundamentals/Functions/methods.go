package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("New Name: ", g.name)

	g.greetReceiver()
	fmt.Println("New Name: ", g.name)

}

type greeter struct {
	greeting string
	name     string
}

// this has a param before the func name
// this is what makes the function turn into a method.
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Tom" // since we are copying it to the variable and not the greeter objexct itself
	// it will not be printed as "Tom" in the main func
}

// Receiver
// This is to manipulate the object itself of the underlying data
func (g *greeter) greetReceiver() {
	fmt.Println(g.greeting, g.name)
	g.name = "Tom"
	// since it is manipulating the object itself, it will be printed as
	// "Tom" in the main func
}
