// Program to pass pointer as a function argument

package main

import "fmt"

// function definition with a pointer argument
func update(num *[]int) {

	// dereferencing the pointer
	s := append(*num, 5)
	fmt.Println(s)

}

func main() {

	var number = 55
	my_slice := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice:", my_slice)

	// function call
	update(&my_slice)

	fmt.Println("The number is", number)

}
