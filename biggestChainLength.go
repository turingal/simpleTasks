package main

import (
	"fmt"
	"regexp"
)

func main() {
	var bytes string

	fmt.Scanf("%s", &bytes)
	re := regexp.MustCompile("1+|0+")

	fmt.Println(getLength(re.FindAllString(bytes, -1)))
}

func getLength(chains []string) (res int) {
	length := 0

	for i := 0; i <= len(chains)-1; i++ {
		if len(chains[i]) > length {
			length = len(chains[i])
		}
	}

	return length
}
