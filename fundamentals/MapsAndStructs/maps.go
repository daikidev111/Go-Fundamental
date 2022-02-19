package main

import "fmt"

func main() {
	// map[key_type]value_type
	statePopulations := map[string]int{
		"California": 1,
		"Texas":      2,
		"Florida":    3,
	}
	fmt.Println(statePopulations)

	//another way of creating a map
	another_map := make(map[string]int)
	fmt.Println(another_map)

	// delete operation
	delete(statePopulations, "Florida")
	fmt.Println(statePopulations)

	//assign a new value to a spec key
	statePopulations["Texas"] = 3
	fmt.Println(statePopulations)

	// what happens if you print out the key that does not exist?
	fmt.Println(statePopulations["Florida"]) // it returns 0 !!!

	// how to check if it is a valid key existing in the map?
	pop, ok := statePopulations["Florida"]
	fmt.Println(pop, ok) // it returns 0, false because it is not found in the map
	// so you can use ok var

	sp := statePopulations
	delete(statePopulations, "Texas")
	fmt.Println(sp, "===", statePopulations)
	// REMEMBER THAT IT AFFECTS THE ASSOGNED MAP/ARR/SLICE!!

	// =========================EOF============================

}
