/*
Author: Stephen Dunne
Date: 28.10.2021
Title: Slices
Desc: Demonstrate Slices functionality in Go
Characteristics of a Slice in Go:
-Not a fixed Size
-Sequence
-Contents are of the same type
*/

package main

import "fmt"

func main() {
	var x []int         // A slice declaration
	y := make([]int, 5) // Using the make function to create a slice. y is based on an array of length 5. IT can be smaller than 5.
	z := make([]int, 5, 10)
	arr := [5]int{1, 2, 3, 4, 5}
	a := arr[0:4]
	// x[4] = 10 // This will give an error
	y[2] = 60
	z[4] = 55
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(a)

	// Appending example
	sl1 := []int{1, 2, 3}
	sl2 := append(sl1, 4, 5)
	fmt.Println(sl1, sl2)

	// Copy example
	sl3 := []int{1, 2, 3}
	sl4 := make([]int, 2)
	copy(sl4, sl3) // Copy (destination, source) == Copy the contents of sl3 into sl4
	fmt.Println(sl3, sl4)
}
