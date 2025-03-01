package main

import "fmt"

func main(){
	fmt.Println("--- Maps in Go ---")

	languages := make(map[string]string)

	languages["JS"]="JavaScript"
	languages["PY"]="Python"
	languages["GO"]="Golang"
	languages["RB"]="Ruby"

	fmt.Println("List of programing Languages: ", languages)
	fmt.Println("JS is short for: ", languages["JS"])
	
	// Deleting an element from map
	delete(languages, "RB")
	fmt.Println("List of programing Languages: ", languages)

	// looping through map
	for key, value := range languages{
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}