package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	fmt.Println("Adder is:", adder(3, 5))
	sum, desc := sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(desc, sum)
}

func greeter() {
	fmt.Println("Chaliye shuru karte hai")
}

func adder(a int, b int) int {
	return a + b
}

func sum(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "The Summation is:"
}
