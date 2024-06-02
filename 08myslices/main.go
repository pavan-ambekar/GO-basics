package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var superHeros = []string{"ironman", "batman", "spiderman"}
	fmt.Printf("Type of superHeros is: %T\n", superHeros)

	superHeros = append(superHeros, "superman", "falcon")
	fmt.Println(superHeros)

	// superHeros = superHeros[1:]
	// fmt.Println(superHeros)

	superHeros = superHeros[1:3]
	fmt.Println(superHeros)

	scores := make([]int, 4)

	scores[0] = 89
	scores[1] = 67
	scores[2] = 45
	scores[3] = 90
	// scores[4] = 87

	scores = append(scores, 34,56,78)
	fmt.Println(scores, len(scores))
	
	fmt.Println(sort.IntsAreSorted(scores)) 
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	// Remove a value to slice by index
	var languages = []string{"English","Kannada","Telugu","Hindi","Tamil"}
	fmt.Println(languages)
	index := 2
	languages = append(languages[:index], languages[index+1:]...)
	fmt.Println(languages)
} 
