package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)	// We are using a walrus operator to declare and initialize the variable as we don't know what type of data we might be using.
	fmt.Println("Enter your age:")
	
	// comma ok || error
	// As there is no try-catch in Golang, we use this syntax.

	// We use an underscore if we don't want to deal with the error and vice versa
	input, _ := reader.ReadString('\n')
	
	fmt.Println("Thank you for providing your age: ", input)
}