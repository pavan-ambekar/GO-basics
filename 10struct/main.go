package main

import "fmt"

func main() {
	fmt.Println("Structs in GO")
	// No Inheritance in GO; No super or parent

	newUser := User{"Pavan", "pavan@gmail.com", false, 24}
	fmt.Println(newUser)
	fmt.Printf("userDetails are: %+v\n", newUser)
	fmt.Printf("Name: %v, Email: %v\n", newUser.Name, newUser.Email)
}

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}
