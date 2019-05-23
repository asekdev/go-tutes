package main

import (
	"fmt"
)

func main() {
	bestFinish := findBestFinish(10, 2, 5, 6, 44, 333, 21)

	fmt.Println("Best finish is =", bestFinish)
}

func findBestFinish(finishes ...int) int {

	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
