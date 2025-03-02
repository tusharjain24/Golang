package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkError(err error){
	if err != nil {
		panic(err)
	}
}

func main(){
	fmt.Println("--- Web Requests in Go ---")

	const url = "https://jsonplaceholder.typicode.com/posts"

	response, error := http.Get(url)

	checkError(error)
	fmt.Printf("Response is of type: %T\n", response)
	// fmt.Printf("Response: %+v\n", response)
	
	defer response.Body.Close()	//Caller's responsibility to close the response body

	dataBytes, error := io.ReadAll(response.Body)
	checkError(error)

	content := string(dataBytes)
	fmt.Println("content: ", content)
}