/*
Author: Stephen Dunne
Date: 30.07.2020
Title: HelloWorld
*/
package main

import "fmt" // fmt = format == stdio in C

//This is a comment, just like in C
func main() {
	x := 2       // Short form of var x int = 2
	y := "Hello" // Short form of var y string = "Hello"
	var (        // Alternative method of declaring multiple variables at once
		a = 10
		b = 11
		c = 12
	)
	const num int = 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(num)
	fmt.Println(a, b, c)

	// var name string = "Stephen"
	// var age int = 22
	// fmt.Println("Hello, my name is "+name+" I am %d ", age)
}
