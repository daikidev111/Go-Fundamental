package main

import "fmt"

func main() {
	var a int = 42

	// How to assign a pointer
	var b *int = &a    // declaring that b is a pointer to the a
	fmt.Println(&a, b) // 0xc000014138 0xc000014138 it is holding the same address

	// how to de-reference the pointer
	fmt.Println(*b) // this prints out 42

	// Why do we use pointers?
	// We use pointer to keep the data consistency across different vars

	// you can even dereference the pointer and use that to change the value
	*b = 14
	fmt.Println(a) // this wil return 14

	// THERE IS NO POINTER ARITHMETIC!!!

	// you can intialize struct like this
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// another way of initialization
	// this will return &{0}
	var ms2 *myStruct
	ms2 = new(myStruct)
	fmt.Println(ms2)

	// what value does the pointer to the struct hold?
	// the pointer holds nil
	// Best practice is to check if the pointer is a nil poitner if you are accepting
	// pointers as arguments
	var ms3 *myStruct
	fmt.Println(ms3) // <nil>
	ms3 = new(myStruct)
	fmt.Println(ms3)

	// it does not need to deference like (*ms4).foo = 100.
	// Thanks to the compiler, you can access to the field without dereferencing
	var ms4 *myStruct
	ms4 = new(myStruct)
	ms4.foo = 100
	fmt.Println(ms4)

}

type myStruct struct {
	foo int
}
