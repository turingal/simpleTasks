package main

import (
	"fmt"
	"sort"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)
	counter := make(map[int]int)

	partOfDivision := number

	for partOfDivision != 0 {
		counter[partOfDivision%10]++
		partOfDivision = partOfDivision / 10
	}

	printCounter(counter)
}

func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
