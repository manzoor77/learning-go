// Golang program to reverse a string
package main

// importing fmt
import "fmt"

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {

	// Reversing the string.
	str := "Geeks"

	// returns the reversed string.
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
