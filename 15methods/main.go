package main

import "fmt"

func main() {
	fmt.Println("Methods in GO")
	// No Inheritance in GO; No super or parent

	newUser := User{"Pavan", "pavan@gmail.com", false, 24}
	fmt.Println(newUser)
	fmt.Printf("userDetails are: %+v\n", newUser)
	fmt.Printf("Name: %v, Email: %v\n", newUser.Name, newUser.Email)
	newUser.GetStatus()
	newUser.NewMail()
	fmt.Printf("userDetails are: %+v\n", newUser)
}

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}

func (u User) GetStatus() {
	fmt.Println("Is user verified:", u.IsVerified)
}

func (u User) NewMail() {
	u.Email = "new@gmail.com"
	fmt.Println("Email is:", u.Email)
}
