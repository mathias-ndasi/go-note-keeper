package main

import "fmt"

func main() {
	fmt.Println("Sum results: ", sumUp(1, 2, 3, 4, 7))
}

func sumUp(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
