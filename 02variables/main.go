/*

* In Go, the fmt package provides various formatting verbs that you can use with functions like Printf to format your output. Here are explanations for `%d` and `%v`:

*- **`%d`**: This verb is used to format integers in base 10. It stands for "decimal" and is commonly used when you want to print integer values.

*- **`%v`**: This verb is the "default format" for the value. It is a versatile verb that can be used to print any value in a default format that makes sense for the type of the value. For example, it will print integers as numbers, strings as quoted strings, and slices or arrays in a readable format.

In your code:
```go
fmt.Printf("Printing courses after removing index %d: %v\n", index, courses)
```
- `%d` is used to print the

index

 variable, which is expected to be an integer.
- `%v` is used to print the courses variable, which could be a slice or any other type, and it will be printed in a default, human-readable format.

Here's a breakdown of the line:
- `"Printing courses after removing index %d: %v\n"` is the format string.
-

index is the integer value that will replace `%d`.
-

courses is the value that will replace `%v`.

This will result in an output like:
```
Printing courses after removing index 3: [course1 course2 course4]
```
assuming index is `3` and courses is a slice of strings.
*/

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

package main

import (
	"fmt"
)

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