package main

import "fmt"

func main(){
	fmt.Println("--- Learn about Slices in Golang ---")
	
	// Initialize a Slice
	var fruitList = []string{"Apple", "Banana", "Peach", "Tomato"}
	fmt.Printf("Type of fruit list is %T", fruitList)
	fmt.Println("Fruit list is :", fruitList)

	
}