package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("--- Learn about Slices in Golang ---")
	
	// Initialize a Slice
	var fruitList = []string{"Apple", "Banana", "Peach", "Tomato"}	//If we use this syntax, we need to initialize it as well
	fmt.Printf("Type of fruit list is %T\n", fruitList)
	fmt.Println("Fruit list is :", fruitList)

	// to add new values to fruit list
	fruitList=append(fruitList, "Sapota", "Pomegranate")
	fmt.Println("Fruit list after appending: ", fruitList)

	// We are slicing the fruitlist and now we won't get "apple"
	// Range is non-inclusive
	// fruitList = append(fruitList[1:3])
	fmt.Println("New FruitList after slicing:", fruitList)

	highScores := make([]int, 4)

	// These are default values of the slice
	highScores[0]=234
	highScores[1]=945
	highScores[2]=465
	highScores[3]=867

// If we uncomment it, we will get an error because we have defined it's range to 4 inside make()
	// highScores[4]=777

	// But we can add new values using append method. It will reallocate memory
	highScores = append(highScores, 555, 666, 777)
	fmt.Println("HighScores: ", highScores)
	fmt.Println("Is highScores sorted: ", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("Sorted HighScores: ", highScores)
	fmt.Println("Is highScores sorted: ", sort.IntsAreSorted(highScores))

	// Remove values from slices based on the index.
	var courses = []string{"react.js", "javascript", "golang", "python", "DSA"}
	fmt.Println("Courses: ", courses)

	var index int = 4
	courses=append(courses[:index], courses[index+1:]...)
	fmt.Printf("Printing courses after removing index %d: %v\n", index, courses)
}

// Methods available for slices in Golang are: append, sort, colon[:] (For Slicing)