package main

import "fmt"

func main() {
	fmt.Println("============== 1st =============")

	// basic loops no parenthesises are required
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("============== 2nd =============")

	// two var initialization and increments
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("============== 3rd =============")
	// if i is intialized outside the for loop and want to use it for looping then use ;
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("============== 4th =============")
	// this represnets a while loop
	i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("============== 5th =============")
	// Infinite loop
	// i = 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	fmt.Println("============== 6th =============")

	// See the below code, we want to break the loop once the product result get greater than equal to 3
	for i := 1; i <= 3; i++ {
		for j := 1; i <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break
			}
		}
	}

	// However the result is different.. it retuns
	// 1
	// 2
	// 3
	// 2
	// 4
	// 3
	// it breaks out of the if statement but it continues. So how do we
	// actually break out of the entire loop?
	// ANS: Declare break variable and use it where u want to break out of the inner and outer loop

	fmt.Println("============== 6th SOLUTION =============")
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; i <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	// =========================EOF============================

}
