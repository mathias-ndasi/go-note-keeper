package recursion

import "fmt"

func recursion() {
	number := getUserInput()

	fmt.Println("Factorial Test: ", factorial(number))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	if number < 0 {
		return 0
	}

	return number * factorial(number-1)
}

func getUserInput() int {
	fmt.Print("Enter number to compute factorial: ")

	var number int
	fmt.Scanln(&number)

	return number
}
