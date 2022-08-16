package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	example := "#GoLangCode!$!67"

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(example, "")

	fmt.Printf("A string of %s becomes %s \n", example, processedString)
}
