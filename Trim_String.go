// Go program to illustrate
// how to trim a string
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "!!Welcome to GeeksforGeeks !!"
	str2 := "@@This is the tutorial of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings
	// Using Trim() function
	res1 := strings.TrimSuffix("Manzoor", "Manzoor")
	res2 := strings.Trim(str2, "@$")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
}
