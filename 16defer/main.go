package main

import "fmt"

func main(){
	fmt.Println("--- Defer Statements in Golang ---")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// world, one, two, hello

// Result: Last In First Out
// Hello
// Two
// One
// world

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Result:
// Hello
// 4
// 3
// 2
// 1
// 0
// Two
// One
// World