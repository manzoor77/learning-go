package main

import (
	"fmt"
)

func main() {
	sum := 1
	for sum < 10 {
		sum = sum + sum
	}
	fmt.Println(sum)
}
