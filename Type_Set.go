package main

import "fmt"

// In Type Set we use combination of Both Type Parameter & Inference using Interface
/*type custtype interface {
	float64 | int // Use OR Operator
}*/

func min[T comparable](a T, b T) T {
	if a == b {
		return a
	}
	return b

}
func main() {
	fmt.Println(min(6, 5))
}
