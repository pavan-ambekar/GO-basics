package main

import "fmt"

func main() {
	fmt.Println("Welcome ro Pointers")

	// var ptr *int	
	// fmt.Println("Value of ptr is:",ptr)
	
	myNumber := 13
	var ptr = &myNumber
	fmt.Println("Value of actual ptr is:",ptr)
	fmt.Println("Value of actual ptr is:",*ptr)
	*ptr = *ptr * 2
	fmt.Println("New Value is:", myNumber)
}