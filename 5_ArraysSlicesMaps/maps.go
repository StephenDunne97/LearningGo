/*
Author: Stephen Dunne
Date: 29.10.2021
Title: Maps
Desc: Demonstrate Map functionality in Go
Characteristics of a Map in Go:
-Unordered
-Key-Value pairs
-Aka: Hashtable, Dictionary, Associative Array
*/

package main

import "fmt"

func main() {
	x := make(map[string]string) // x is a map of strings to strings
	x["IE"] = "Ireland"          // Key == IE, Value == Ireland
	x["UK"] = "United Kingdom"
	x["FR"] = "France"
	x["DE"] = "Germany"

	fmt.Println(x["IE"]) // When I search the Key, the Value is returned
	fmt.Println(x["FR"])
	fmt.Println(x["UK"])
	fmt.Println(x["US"]) // When I search for a key that does not exist, nothing is returned
	fmt.Println(x["DE"])
	delete(x, "DE") // Delete a Key:Value pair by inputting the key

	name, ok := x["DE"]   // Check if a Key:Value pair exists
	fmt.Println(name, ok) // This will return False

	name, ok = x["IE"]
	fmt.Println(name, ok) // This will return True

	// Checking if a Key:value pair exists and printing if it does
	if name, ok = x["DE"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("This key:value pair does not exist.")
	}

	if name, ok = x["IE"]; ok {
		fmt.Println("Value", name, "Exists", ok)
	} else {
		fmt.Println("This key:value pair does not exist.")
	}

	y := make(map[int]string) // y is a map of ints to strings
	y[1] = "Stephen Dunne"    // Student Number maps to student name

	fmt.Println(y[1])

	// Map of maps and a shorter way of declaring map contents
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithum",
			"state": "solid",
		},
	}

	fmt.Println(elements["H"]) // This will return a map and its constituent key:value pairs
	// Checking that an element exists, and if it does, printing its key:value pairs
	if el, ok := elements["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
