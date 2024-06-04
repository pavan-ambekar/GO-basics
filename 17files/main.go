package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome tot file in golang")

	content := "This need to go in a file"
	file, err := os.Create("./my_file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./my_file.txt")
}

func readFile(filename string) {
	dataBytes, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Data in file:", string(dataBytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
