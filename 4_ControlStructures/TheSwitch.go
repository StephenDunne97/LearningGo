/*
Author: Stephen Dunne
Date: 28.10.2021
Title: ForLoop
Desc: Demonstrate the Switch statement functionality in Go
*/

package main

import "fmt"

func main() {
	i := 420
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("That is, unfortunately, not recognizable by this program.")
	}
}
