package main

import "fmt"

func main(){
	fmt.Println("--- If Else in Golang ---")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	}else if loginCount > 10 && loginCount < 15 {
		result = "Super User"
	}else{
		result = "Watch out"
	}

	fmt.Println(result)


	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	// Special Syntax in the case of web request handling
	if num := 3; num < 10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is not less than 10")
	}
}