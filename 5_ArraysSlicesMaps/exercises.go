/*
Author: Stephen Dunne
Date: 29.10.2021
Title: Exercises from chapter 5
Desc: Demonstrate competency in understanding of arrays, maps, and slices in Go
*/

package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	// How do you access the fourth element of an array or slice?
	fmt.Println("The 4th element in arr:", arr[3])
	// What is the length of a slice created using make([]int, 3, 9)?
	arr2 := make([]int, 3, 9)
	fmt.Println("The length of arr2:", len(arr2))
	// Given the following array, what would x[2:5] give you?
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5]) // Output: c, d, e. It is inclusive of the min/max indexes
	// Write a program that finds the smallest number in this list:
	y := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	lowest := y[0]
	/*
		for i := 0; i < len(y); i++ {
			if y[i] < lowest {
				lowest = y[i]
			}
		}
	*/

	for _, i := range y {
		if i < lowest {
			lowest = i
		}
	}

	fmt.Println("Lowest:", lowest)

}
