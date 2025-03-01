package main

import "fmt"

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

		viral := User{"Viral", "Jain", "abcd@abc.com", 26, true}

		// %v will print the values of the fields
		fmt.Printf("Details of Tushar are: %v",tushar)

		// %+v will print the field names along with the values
		fmt.Printf("\nDetails of Viral are: %+v\n",viral)

		// %T will print the type of the struct
		fmt.Printf("Type of Tushar is: %T\n",tushar)

		fmt.Printf("Name of Tushar is: %v, %v",tushar.FirstName, tushar.LastName)
}

	// There is no inheritance in Golang
	// There is no concept of classes in Golang such as super, sub, base, derived, etc and parent.

	// Struct is a collection of fields	; As they will be exported, so the name of the struct will be capitalized along with the fields.
type User struct{
	FirstName string
	LastName string
	Email string
	Age int
	Status bool
}