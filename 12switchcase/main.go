package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to LUDO in GO")
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is one you can open")
	case 2:
		fmt.Println("Dice value is 2 spot")
	case 3:
		fmt.Println("Dice value is 3 spot")
	case 4:
		fmt.Println("Dice value is 4 spot")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 spot")
	case 6:
		fmt.Println("Dice value is 6 spot and roll dice again")
	default:
		fmt.Println("What was that")
	}
}
