package main

import "fmt"

func main() {
	var fruits [5]string

	fruits[0] = "apple"

	fmt.Printf("Fruits: %v", fruits)

	// the pointer semantic form of a 'for'
	// fruits[1] will be overwritten!
	for i := range fruits {
		fruits[1] = "pineapple"

		fmt.Println(i)
	}

	// the value semantic form
	// fruits[1] won't be overwritten
	for _, fruit := range fruits {
		fruits[1] = "pineapple"

		fmt.Println(fruit)
	}

	// this is considered bad Go! it's confusing and mixes semantics
	for _, fruit := range &fruits {
		fruits[1] = "pineapple"

		fmt.Println(fruit)
	}

	// this makes a slice with a backing array of 1000 elems
	data := make([]string, 5, 1000)
	fmt.Printf("data: %v", data)

	slice1 := []string{"some", "string"}
	slice2 := make([]string, len(slice1))
	// this creates a new backing array manually
	copy(slice2, slice1)

}
