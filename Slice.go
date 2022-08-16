package main

import "fmt"

func main() {
	// Array Declaration with Type
	/*var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	array1 := [2]int{6, 7}
	fmt.Println(array1)
	// Slice
	fmt.Println(array[:2])*/

	/*names := [4]string{"John","Paul","George","Ringo",}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)*/

	//Lenth and Capicty of Slice
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[2:]
	fmt.Println(s)
}
