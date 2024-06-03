package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in GO")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Tuesday", "Friday", "Saturday"}
	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Println(day)
	}

	rougeValue := 2
	for rougeValue < 20 {
		if rougeValue == 15 {
			break
		}
		if rougeValue == 10 {
			goto jump
		}
		if rougeValue%2 == 1 {
			rougeValue++
			continue
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

jump:
	fmt.Println("Jumped outside")
}
