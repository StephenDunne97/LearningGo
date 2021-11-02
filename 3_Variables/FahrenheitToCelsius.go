package main

import "fmt"

func main() {
	var input float64
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fmt.Scanf("%f", &input)
	output := (input - 32) * 5 / 9
	fmt.Println("In Celsius:", output)
}
