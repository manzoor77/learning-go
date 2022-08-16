package main

import "fmt"

// In Type Inference we have type Constraint
// Its mean Type should be exact same as mentioned in Function type with Call Function
func min[T float64](a T, b T) T {
	return a
}
func main() {
	fmt.Println(min[float64](2, 5))
}
