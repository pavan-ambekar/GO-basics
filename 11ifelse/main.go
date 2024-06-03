package main

import "fmt"

func main() {
	fmt.Println("if else in GO..")
	loginCount := 32

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount < 20 {
		result = "Interesting User"
	} else {
		result = "Super User"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	if num := 5; num < 10 {
		fmt.Println("num is less then 10")
	} else {
		fmt.Println("num is not less then 10")
	}

}
