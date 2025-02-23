package main

import "fmt"

func main(){
	fmt.Println("--- Learn about Arrays in GO ---")

	// Declare an array
	var fruitList [4]string

	// Initialize the array
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "Sapota"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length(Size) of fruit list is:", len(fruitList))

	// Declare and Initialize the variable array

	var vegitableList = [4]string{"Brinjal Onion Potato"}
	fmt.Println("Vegitable list is: ", vegitableList)
	fmt.Println("Length of Vegitable List is:", len(vegitableList))
}