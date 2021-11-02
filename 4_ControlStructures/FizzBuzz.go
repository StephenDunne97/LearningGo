/*
Author: Stephen Dunne
Date: 28.10.2021
Title: ForLoop
Desc:
-Print numbers 1-100
-If the current number is evenly divisble by 3 and 5, print FizzBuzz
-If the current number is only evenly divisble by 3, print Fizz
-If the current number is only evenly divisble by 5, print Buzz
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i += 1 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
