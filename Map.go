package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map is Working")
	/*var a = make(map[string]int)
	a["brand"] = 1
	a["model"] = 2
	a["year"] = 3

	fmt.Printf("%v\n", a)*/
	// Create Map without make Function
	// It will nill by default
	var a = map[string]int{
		"brand": 1,
		"model": 2,
		"year":  3,
	}
	fmt.Printf("%v\n", a)
	// using struct
	type mapstruct struct {
		names string
		id    int
	}
	var b = make(map[int]mapstruct)
	b[1] = mapstruct{
		"manzoor", 5,
	}

	fmt.Println(b[1])
}
