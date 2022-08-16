package main

import "fmt"

var factorial int

func calcFactorial(factorialnum int) int {
	factorial = 1
	if factorialnum < 18 {
		for i := 1; i <= factorialnum; i++ {
			factorial = factorial * i
		}
		return factorial
	} else {
		return 0
	}
}
func main() {

	var factorialnum int

	fmt.Print("Enter any Number to find the Factorial = ")
	fmt.Scanln(&factorialnum)

	factorial = calcFactorial(factorialnum)
	fmt.Println("The Factorial of ", factorialnum, " = ", factorial)
}
