/*
Author: Stephen Dunne
Date: 28.10.2021
Title: ForLoop
Desc: Demonstrate the for loop functionality in Go
*/
package main

import "fmt"

func main() {
	var i int = 1 // i:=1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
	// Go does not use while, do, until, foreach)
	// The below is an alternative way of writing For Loops in Go:
	for i := 1; i <= 10; i++ {
		// Adding addtional functionality to print if the current number is odd or even
		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")
		}
	}
}
