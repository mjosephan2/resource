package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "one,two;three four"
	// Create a regex that matches commas, semicolons, or spaces
	re := regexp.MustCompile(`[,\s;]+`)
	parts := re.Split(input, -1)

	fmt.Println(parts)
}
