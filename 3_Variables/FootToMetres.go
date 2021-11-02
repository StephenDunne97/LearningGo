package main

import "fmt"

func main() {
	var input float64
	metre := 0.3048
	fmt.Print("Enter a measurement in Feet: ")
	fmt.Scanf("%f", &input)
	output := metre * input
	fmt.Println("In Metres:", output)
}
