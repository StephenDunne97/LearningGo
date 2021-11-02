/*
Author: Stephen Dunne
Date: 28.10.2021
Title: ForLoop
Desc: Refresher
*/
package main

import "fmt"

func main() {
	var x int = 22
	y := 44
	name := "Stephen"
	var inputName string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &inputName)
	fmt.Println("This is the name you entered: ", inputName)
	fmt.Print("X + Y = ", x+y, " This is my name: ", name)
	fmt.Print("This is a test to see if this appears on a new line")
	fmt.Println("The above will not appear on a new line, nor will this :(")
	fmt.Print("This will though")
}
