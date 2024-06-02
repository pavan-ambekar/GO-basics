package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	var superHeros [4]string
	superHeros[0] = "IronMan"
	superHeros[1] = "SuperMan"
	superHeros[3] = "BatMan"
	fmt.Println("SuperHeros:", superHeros)
	fmt.Println("Total SuperHeros:", len(superHeros))

	var days = [3]string{"Monday","Tuesday","Wednesday"}
	fmt.Println("Days:", days)
	fmt.Println("TotalDays:", len(days))

}