package main

import (
	"fmt"
)

func main() {
	var bytes string

	fmt.Scanf("%s", &bytes)
	fmt.Println(getLength(bytes))
}

func getLength(stringData string) (res int) {
	if len(stringData) == 0 {
		return 0
	}

	lengthCurr := 1
	lengthMax := 0

	for i := 1; i < len(stringData); i++ {
		if stringData[i] == stringData[i-1] {
			lengthCurr++
			continue
		}

		if lengthCurr > lengthMax {
			lengthMax = lengthCurr
			lengthCurr = 1
		}
	}

	if lengthCurr > lengthMax {
		return lengthCurr
	}

	return lengthMax
}
