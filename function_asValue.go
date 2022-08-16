package main

import (
	"fmt"
	"math"
)

// square function
func square(val float64) float64 {
	return (math.Pow(val, 2)) //Pow(2,2)
}

// cube function
func cube(sqr func(float64) float64) float64 { //sqr==>square
	return math.Pow(sqr(2), 3) //square(2)
} //pow(4,2)
func main() {
	fmt.Println(cube(square)) // pass reference of square to cube function
}
