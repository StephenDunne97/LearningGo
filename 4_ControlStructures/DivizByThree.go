/*
Author: Stephen Dunne
Date: 28.10.2021
Title: DivizByThree
Description: Print all numbers between 1-100 that are evenly divisble by 3.
*/

package main

import "fmt"

func main() {
	for i := 0; i <= 100; i += 1 {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
