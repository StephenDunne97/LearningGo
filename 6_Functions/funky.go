/*
Author: Stephen Dunne
Date: 04.11.2021
Title: Functions
Desc: Demonstrate functions functionality in Go
*/

package main

import "fmt"

func average(scores []int) int {
	total := 0
	average := 0

	for _, v := range scores {
		total += v
	}
	average = total / len(scores)
	return average
}

// Functions can return multiple values
func f1() (int, int) {
	return 5, 9
}

// Variadic function
func variadic(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	scores := []int{98, 93, 77, 82, 83}
	fmt.Println(average(scores))
	x, y := f1() // Assigning multiple values from a function
	fmt.Println(x, y)
	fmt.Println(variadic(1, 2, 3))   // Using a variadic function
	fmt.Println(variadic(scores...)) // Passing a slice to a variadic function

	// A local function
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println("Add fxn output:", add(1, 3))

	// Demo local function scope
	z := 0
	increment := func() int {
		z++ // This function can access all variables contained within main
		return z
	}
	fmt.Println("Increment fxn first run:", increment())
	fmt.Println("Increment fxn second run:", increment())

}
