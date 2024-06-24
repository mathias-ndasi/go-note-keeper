package functions

import "fmt"

type TransformFunction func(int) int

func function() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	doubleTransformedNumbers := transformNumbers(&numbers, doubleNumbers)
	tripleTransformedNumbers := transformNumbers(&numbers, tripleNumbers)

	fmt.Println(numbers)
	fmt.Println(doubleTransformedNumbers)
	fmt.Println(tripleTransformedNumbers)
	fmt.Println(transformNumbers(&[]int{1, 2, 3}, getTransformerFunction(&[]int{1, 2, 3})))
}

func getTransformerFunction(numbers *[]int) TransformFunction {
	if (*numbers)[0] == 0 {
		return doubleNumbers
	}

	return tripleNumbers
}

func transformNumbers(numbers *[]int, transformer TransformFunction) []int {
	transformedNumbers := []int{}

	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transformer(number))
	}

	return transformedNumbers
}

func doubleNumbers(number int) int {
	return number * 2
}

func tripleNumbers(number int) int {
	return number * 3
}
