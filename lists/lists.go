package lists

import (
	"fmt"
	"time"
)

type Product struct {
	id    int64   `json:"id"`
	title string  `json:"title"`
	price float64 `json:"price"`
}

func NewProduct(title string, price float64) *Product {
	return &Product{
		id:    time.Now().UnixNano(),
		title: title,
		price: price,
	}
}

func list() {
	hobbies := [3]string{"Video gaming", "Switching", "Coding"}
	fmt.Println(hobbies)
	fmt.Println("First hobby: ", hobbies[0])

	firstHobbySlices := hobbies[1:]
	fmt.Println("Second and third hobbies: ", firstHobbySlices)

	secondHobbySlices := hobbies[:2]
	thirdHobbySlices := hobbies[0:2]

	fourthHobbySlices := secondHobbySlices[1:3]

	courseGoals := []string{"Learn", "Work"}
	courseGoals[1] = "Home"
	courseGoals = append(courseGoals, "Money")

	fmt.Println(thirdHobbySlices)
	fmt.Println(fourthHobbySlices)
	fmt.Println(courseGoals)

	products := []Product{*NewProduct("Hi", 20.69), *NewProduct("Hello", 30.6)}
	products = append(products, *NewProduct("Test", 40.69))

	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
