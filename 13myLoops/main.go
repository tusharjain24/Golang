package main

import "fmt"

func main(){
	fmt.Println("--- Loops in Golang ---")

	weekDays := []string{"Monday", "Tuesday", "Wednesday", "Thrusday", "Friday"}
	fmt.Println("Week Days: ", weekDays)

	// For Loop
/*	for d:= 0; d < len(weekDays) ; d++{
		fmt.Println("Day:", weekDays[d])
	}
*/

/*
	for i := range weekDays{
		fmt.Println("day: ", weekDays[i])
	}
*/

/*
	for index, day := range weekDays{
		fmt.Printf("Index is %v and the Day is %v\n", index, day)
	}
*/

	lco:
		fmt.Println("Jumping at golang.org to learn GO language")
	rogueValue := 1
	for rogueValue < 10{

		if rogueValue == 5{
			// break
			rogueValue++
			continue	// skips the value. Always remember to change the value of the variable
		}

		if rogueValue == 2{
			goto lco
		}

		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}


}