package main

import (
	"fmt"
	"net/url"
)

const myUrl = "http://localhost:8080/learn?course=go&module=web"

func checkError(err error){
	if err != nil{
		panic(err)
	}
}

func main(){
	fmt.Println("--- Handling URLs in Go ---")
	fmt.Println("URL:", myUrl)

	result, err := url.Parse(myUrl)
	checkError(err)

	fmt.Printf("Type of result: %T\n", result)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Port:", result.Port())
	fmt.Println("Path:", result.Path)
	fmt.Println("Query(Params):", result.RawQuery)

	// Better way to get the query params
	params := result.Query()
	fmt.Printf("Type of params: %T\n", params)

	for key, value := range params{
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Constructing a URL
	newUrl := &url.URL{
		Scheme: "https",	// This is mandatory
		Host: "jsonplaceholder.typicode.com",	// This is mandatory
		Path: "/posts",	// This is optional
		RawQuery: "course=go&module=web",	// This is optional
	}
	// https://jsonplaceholder.typicode.com/posts

	anotherUrl := newUrl.String()
	fmt.Println("Another URL:", anotherUrl)
}