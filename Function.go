package main

import (
	"fmt"
)

func main() {

	fmt.Println(testfunction(1, 2))
	//Call Swap Function
	FirstName, LastName := swap("Manzoor", "Faisal")
	fmt.Println(FirstName, LastName)

}
func testfunction(x int, y int) int {
	fmt.Println("I am Function")
	var z = x + y
	return z

}

// Swap Using Function
func swap(x, y string) (string, string) {
	return y, x
}
