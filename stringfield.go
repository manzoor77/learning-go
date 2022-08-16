package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I am learnig Golang"
	sli := strings.Fields(str)
	fmt.Println(sli[0])
	fmt.Println(sli[2])
	fmt.Println(str)
}
