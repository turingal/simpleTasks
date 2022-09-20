package main

import (
	"fmt"
)

func main() {
	var text string
	var width int

	fmt.Scanf("%s %d", &text, &width)

	if len(text) <= width {
		fmt.Printf(text)

		return
	}

	for i := 0; i < width; i++ {
		fmt.Printf(string(text[i]))
	}

	fmt.Println("...")

}
