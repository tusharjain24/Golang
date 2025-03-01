package main

import "fmt"

func main(){
	fmt.Println("--- Functions in Golang ---")
	greeting()

	result := adder(3,5)
	fmt.Println("Result of addition is: ", result)

	proResult, _ := proAdder(4, 5, 6, 7, 8)
	fmt.Println("Result of proResult is: ", proResult)


	// Function inside another function is not allowed in golang. We have to create the function outside and then call it inside

	// Anonymous function and IIFE's also exist in Golang
/*
	func greeterTwo(){
		fmt.Println("Another method")
	}
	greeterTwo()
*/

}
// Syntax of function
/*

func functionName(param1 dataType, param2 dataType) Signature(dataType that the function returns){

}

*/

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}


func proAdder(values ...int) (int, string){
	total := 0

	for _, value := range values{
		total += value
	}
	return total, "Hi result of proAdder"
}

func greeterTwo(){
	fmt.Println("Another method")
}

func greeting(){
	fmt.Println("Namaste from Golang")
}