package main

import "fmt"

// authToken := "dsakf" // this will throw error walrus operator works only inside func
var authToken = "new token"

const LoginToken = "token" // Public 

func main() {
	var username string = "Pavan"
	fmt.Println(username)
	fmt.Printf("Variable iss type of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable iss type of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable iss type of type %T \n", smallVal)

	var smallFloat float32 = 255.045095996
	fmt.Println(smallFloat)
	fmt.Printf("Variable iss type of type %T \n", smallFloat)

	var longFloat float64 = 255.045095996
	fmt.Println(longFloat)
	fmt.Printf("Variable iss type of type %T \n", longFloat)

	// Default values and some aliases
	var intVal int
	fmt.Println(intVal)
	fmt.Printf("Variable iss type of type %T \n", intVal)

	// Implicit type
	var website = "google.com"
	fmt.Println(website)
	// website=12 //THIS WILL THROW THE ERROR

	// no var style
	numberOfUser := 4500.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable iss type of type %T \n", LoginToken)
}
