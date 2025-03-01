// In Go, both methods and functions are used to perform actions, but they have some key differences:
//
// Functions:
// - Functions are defined outside of any type.
// - They can be called independently without any receiver.
// - Syntax: func functionName(parameters) returnType { ... }
//
// Methods:
// - Methods are functions with a special receiver argument.
// - The receiver can be either a value or a pointer of a named type.
// - Methods are associated with the type and can be called on instances of that type.
// - Syntax: func (receiverType receiverName) methodName(parameters) returnType { ... }
//
// Example of a function:
// func add(a int, b int) int {
//     return a + b
// }
//
// Example of a method:
// type Rectangle struct {
//     width, height int
// }
//
// func (r Rectangle) area() int {
//     return r.width * r.height
// }
//
// Explanation:
// - The `add` function takes two integers as parameters and returns their sum. It is a standalone function and can be called as `add(3, 4)`.
// - The `area` method is associated with the `Rectangle` type. It takes a `Rectangle` instance as its receiver and returns the area of the rectangle. It can be called on a `Rectangle` instance like `rect.area()`, where `rect` is an instance of `Rectangle`.

package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct{
	FirstName string
	LastName string
	Email string
	Age int
	Status bool
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the new email: ")

	data, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error", err)
	}

	u.Email = data
	fmt.Println("Email of the user has been updated: ", u.Email)
}

func main(){
	fmt.Println("--- Structs in Golang ---")
	
	//using the below struct, we can create multiple users
	// User is a struct
	
	tushar := User{
		FirstName: "Tushar",
		LastName: "Jain",
		Email: "abc@gmail.com",
		Age: 25,
		Status: true,
	}
	fmt.Println("Details of Tushar: ", tushar)
	
	viral := User{"Viral", "Jain", "abcd@abc.com", 26, true}
	viral.GetStatus()
	
	// The Method in Go takes a copy of the struct so whatever manipulation we do to the object it won't change original object 
	fmt.Println("Calling NewEmail function")
	viral.NewEmail()
	fmt.Printf("Details of Viral: %+v\n", viral)
}