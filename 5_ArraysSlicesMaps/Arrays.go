/*
Author: Stephen Dunne
Date: 28.10.2021
Title: Arrays
Desc: Demonstrate Array functionality in Go
Characteristics of an Array in Go:
-Fixed Size 
-Sequence 
-Contents are of the same type
*/
package main

import "fmt"

func main() {
	var x [5]int
	total := 0
	example := [5]int{
		2345,
		59845,
		3285,
		587,
		4892,
	}

	x[1] = 5212

	fmt.Println(x)    // Print the entire array
	fmt.Println(x[1]) // Print the value at the specified index

	for i := 0; i < len(example); i++ { // Loop through the array and add each value total
		total += example[i]
	} // End for
	fmt.Println("Average:", total/len(example)) // Print the average of all values in 'example' array

	// A shorter way to write the above:
	total = 0
	for _, value := range example {
		total += value
	} // End for
	fmt.Println("Average:", total/len(example))
}
