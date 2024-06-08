package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://guthib.com/"

func main() {
	fmt.Println("Web Request")

	res, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Res is type of: %T\n", res)
	fmt.Println("Status Code:", res.StatusCode, res.Status)

	defer res.Body.Close() // callers responsibility to close the connection

	databytes, err := io.ReadAll(res.Body)
	checkNilErr(err)

	content := string(databytes)
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
