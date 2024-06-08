package main

import (
	"fmt"
	"net/url"
)

const my_url = "https://pavan.dev:8000/projects?language=js&freamework=react"

func main() {
	fmt.Println("URL handling in golang")
	fmt.Println(my_url)

	// parsing the URL
	result, err := url.Parse(my_url)
	checkNilErr(err)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("host:", result.Host)
	fmt.Println("path:", result.Path)
	fmt.Println("port", result.Port())
	fmt.Println("query", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("type of query params:%T\n", qparams)

	fmt.Println("language:", qparams["language"])

	for key, val := range qparams {
		fmt.Println(key, ":", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "pavan.dev",
		Path:    "projects",
		RawPath: "language=js",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
