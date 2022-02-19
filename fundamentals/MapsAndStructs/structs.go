package main

import "fmt"

// struct gathers related fields to represent one concept
// one of the adv of this is that it can collect mix data types,
// which provides the flex to this data structure, compared to the other, such as array
type Doctor struct {
	// list of fields
	number    int
	actorName string
	companies []string
}

// How to achieve inheritance in go?
// => we can use composition instead through embedding

type Animal struct {
	Name   string
	Origin string
}

// type Bird struct { // this is not a dependent struct with Animal
// 	speedKPH float32
// 	CanFly   bool
// }

// This is how you do embedding ...

type Bird struct {
	Animal   // animal struct embedded!
	SpeedKPH float32
	CanFly   bool
}

// see the below code where it instantiates bird

func main() {

	aDoctor := Doctor{
		number:    3,
		actorName: "daiki",
		companies: []string{"a", "b"},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName) // use . to access to each field

	// anonymous struct
	// condensed to one line, but it cannot be used anywhere else since it is anonymous
	// usecase: when you get a information that has to be organized for one and only one specific situation
	bDoctor := struct{ name string }{name: "John K"}
	fmt.Println(bDoctor)

	// returns
	// 	bDoctor {John K}
	// another_bDctor {Tom F}
	// this shows that it is pointed to an independent data set
	another_bDoctor := bDoctor
	another_bDoctor.name = "Tom F"
	fmt.Println("bDoctor", bDoctor)
	fmt.Println("another_bDctor", another_bDoctor)

	// returns
	// 	bDoctor {Tom F}
	// another_bDoctor2 &{Tom F}
	// another_bDoctor2 is pointed to the dataset of bDoctor and hence its memory is overwritten
	another_bDoctor2 := &bDoctor
	another_bDoctor2.name = "Tom F"
	fmt.Println("bDoctor", bDoctor)
	fmt.Println("another_bDoctor2", another_bDoctor2)

	b := Bird{
		Animal:   Animal{Name: "Kiwi", Origin: "New Zealand"},
		SpeedKPH: 10,
		CanFly:   false,
	}
	fmt.Println(b)

	// From the above, it shows that it is a composition relationship
	// It is a "has-a" relationship and they cannot be used interchangeably
	// to use data interchangeably we can use interfaces

	// embedding usecase: get some base behaviour into a custom type

	// =========================EOF============================

}
