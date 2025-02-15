package main

import (
	"fmt"
)

/*

* '%T' is a placeholder to find the type of a variable.
* 'Printf()' is a statement where we can write formatted string but we cannot write formatted string in Println() statement.
* For both types int and float the init value of their value is '0'.
* For String type variables the init value is an empty string('')
* For Bool type variables the init value is 'false'.

* Note:	The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
 */

/*
* A numeric constant has no type until it’s given one, such as by an explicit conversion.

	fmt.Println(int64(d))

* A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
*/

const LoginToken string = "qwertyuiop12345g"
func main(){
	fmt.Println("--- STRING VARIABLE ---")
	var initStringValue string
	fmt.Printf("Initial value when we initialize a variable with type %T is: \n", initStringValue)
	fmt.Println(initStringValue)
	var username string = "Tushar"
	fmt.Printf("username variable is of type: %T and it value is\n", username)
	fmt.Println(username)

	fmt.Println("--- BOOL VARIABLE ---")
	var initBoolValue bool
	fmt.Printf("Initial value when we initialize a variable with type %T is: \n", initBoolValue)
	fmt.Println(initBoolValue)
	var isLoggedIn bool = false
	fmt.Printf("isLoggedIn variable is of type: %T \n", isLoggedIn)

	fmt.Println("--- INT VARIABLE ---")
	var initValue int
	fmt.Printf("Initial value when we initialize a variable with type %T is: \n", initValue)
	fmt.Println(initValue)
	fmt.Println("uint8 Range: 0 to 255")
	fmt.Println("uint32 Range: 0 to 4294967295")
	fmt.Println("int8 Range:  to ")
	var smallValue uint8 = 255
	fmt.Printf("username variable is of type: %T \n", smallValue)
	fmt.Println(smallValue)
	
	fmt.Println("--- Float Variable ---")
	var initFloat32Value float32
	fmt.Printf("Initial value when we initialize a variable with 'float32' type is %T\n", initFloat32Value)
	fmt.Println(initFloat32Value)
	var initFloat64Value float64
	fmt.Printf("Initial value when we initialize a variable with type %T is:\n", initFloat64Value)
	fmt.Println(initFloat64Value)

	fmt.Println("--- Different ways to declare variables ---")
	
	// Implicit type
	var language = "Golang"
	fmt.Print("Declared an Implicit type variable: ", language)

	// no var style or shorthand
	numberOfUsers := 30000	// := is known as Walrus operator
	fmt.Print("Number of users", numberOfUsers)

}