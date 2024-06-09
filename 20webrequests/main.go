package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "http://localhost:8000"

func main() {
	fmt.Println("web requests in golang")
	PerformGetRequest(myurl + "/get")
	PerformPostJsonRequest(myurl + "/post")
	PerformPostFormRequest(myurl + "/postform")
}

func PerformGetRequest(myurl string) {
	fmt.Println(myurl)

	res, err := http.Get(myurl)
	CheckNilErr(err)
	defer res.Body.Close()

	fmt.Println("Status Code:", res.StatusCode)
	fmt.Println("Content length:", res.ContentLength)

	var responseString strings.Builder

	content, err := io.ReadAll(res.Body)
	CheckNilErr(err)

	byteCount, err := responseString.Write(content)
	CheckNilErr(err)

	fmt.Println("byte Count:", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest(myurl string) {
	//  fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"Pavan A",
			"age":25
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	CheckNilErr(err)
	defer res.Body.Close()

	var responseBuilder strings.Builder
	fmt.Println("Status Code:", res.StatusCode)

	content, err := io.ReadAll(res.Body)
	CheckNilErr(err)

	_, err = responseBuilder.Write(content)
	CheckNilErr(err)

	fmt.Println(responseBuilder.String())
}

func PerformPostFormRequest(myurl string) {
	// form data

	data := url.Values{}
	data.Add("firstname", "pavan")
	data.Add("lastname", "A")
	data.Add("email", "pavan@gmail.com")

	res, err := http.PostForm(myurl, data)
	CheckNilErr(err)
	defer res.Body.Close()

	var resBuilder strings.Builder
	fmt.Println("Status code:", res.StatusCode)
	content, err := io.ReadAll(res.Body)
	CheckNilErr(err)

	_, err = resBuilder.Write(content)
	CheckNilErr(err)
	fmt.Println(resBuilder.String())
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
