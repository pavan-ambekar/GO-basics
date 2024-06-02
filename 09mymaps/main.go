package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["TS"] = "Typescript"

	fmt.Println("List all languages:", languages)
	fmt.Println("JS:", languages["JS"])

	delete(languages, "TS")
	fmt.Println("List all languages:", languages)

	// iterating
	for key, value := range languages {
		fmt.Printf("%v : %v\n", key, value)
	}
}
