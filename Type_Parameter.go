package main

import "fmt"

// use Any Keyword to accept any type from user
func min[T any](a T, b T) T {

	return a

}
func main() {
	fmt.Println(min(2, 5))
	fmt.Println(min[int](2, 5))
	fmt.Println(min[float64](2, 5))

}
