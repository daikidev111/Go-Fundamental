package main

import (
	"fmt"
	"net/http"
)

// In go, there is no exeption methods
// panic is defined as a situation where it cannot continue to function
func main() {

	// this throws panic because int cannot divide by zero
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	// this will throw a panic along with the message
	fmt.Println("Start")
	panic("somrthing bad happened")
	fmt.Println("End")

	// Practical Example
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}
