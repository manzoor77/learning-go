// Go program to illustrate how
// to iterate over the string
// using for range loop
package main

import "fmt"

// Main function
func main() {

	// String as a range in the for loop
	for index, s := range "GeeksForGeeKs" {

		fmt.Printf("The index number of %c is %d\n", s, index)
	}

}
