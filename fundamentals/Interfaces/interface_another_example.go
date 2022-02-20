package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

}

type Incrementer interface {
	Increment() int
}

type IntCounter int // it does not have to be a struct like the other example

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
