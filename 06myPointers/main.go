package main

import "fmt"

func main(){
	fmt.Println("--- Learn about Pointers ---")

	// *int tell the compiler that the ptr variable is a pointer that stores integer value.
	var ptr *int
	fmt.Println("Default value of pointer is: ", ptr)

	// &VariableName. & is used for reference to the Variable
	myNumber:= 13
	var myNumberPointer = &myNumber
	
	// Will the display the actual address 
	fmt.Println("Value of myNumberPointer is: ", myNumberPointer)
	fmt.Println("Value stored inside the address stored in the myNumberPointer is: ", *myNumberPointer)

	// Manipulation of value stored inside the address of the pointer
	var newValue int = *myNumberPointer * 2
	fmt.Println("new Value: ", newValue)
}