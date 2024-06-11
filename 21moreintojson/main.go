package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React JS", 299, "Web", "123123123", []string{"web-dev", "js"}},
		{"Mongo DB", 399, "Web", "111111111", []string{"database", "nosql"}},
		{"HTML", 199, "Web", "222222222", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	CheckNilError(err)

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"courseName": "React JS",
			"Price": 299,
			"website": "Web",
			"tags": [ "web-dev", "js"]
		}
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// Some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Println(k, ":", v)
	}

}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
