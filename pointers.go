/*
Author: Stephen Dunne
Date: 04.11.2021
Title: Pointers
Desc: Demonstrate Pointers functionality in Go
Pointers refer to a location in memory.
When you alter the value of a pointer, you change the value stored in that location, rather than the value of the variable.

& is used to get the address of a variable in memory
e.g: &x gets x's address in memory
*xPtr is used to point to the address obtained by using &x
*xPtr = 5 means; Put 5 into the location that is being pointed to

*/

package main

import "fmt"

func pointerDemo(x int, yPtr *int) {
	x = 5
	*yPtr = 5 // Put 5 into the address that yPtr points to
}

func main() {
	x := 0
	y := 0             // Could delcare a variable like "z := &y and pass that to the function."
	pointerDemo(x, &y) // & gets the address in memory where y is stored
	fmt.Println(x, y)  // X stays as 0, Y changes to 5
}
