/*
Author: Stephen Dunne
Date: 30.07.2020
Title: Operands in Go
*/

/*
Operands support by Go:
+ - Addition
- - Subtraction
* - Multiplication
/ - Division
% - Remainder
*/
package main

import "fmt"

func main() {
	//Operand demo
	fmt.Println("1+1=", 1+1)
	fmt.Println("2*2=", 2*2)
	fmt.Println("1-1=", 1-1)
	fmt.Println("4/2=", 4/2)
	fmt.Println("11%5=", 11%5)

	//String demo
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")

	//Boolean demo
	fmt.Println(true && true)  //AND
	fmt.Println(true && false) //AND
	fmt.Println(true || false) // OR
	fmt.Println(!true)         //NOT
}
