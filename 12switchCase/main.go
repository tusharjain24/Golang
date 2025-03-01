package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("--- Switch Case in Golang ---")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number is: ", diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can move two spots")
	case 3:
		fmt.Println("Dice value is 3 and you can move three spots")
	case 4:
		fmt.Println("Dice value is 4 and you can move four spots")
			// explicitly mention fallthrough if we have a usecase where we want to execute the lines of code after this case.
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 and you can move five spots")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6 and you can move six spots")
	default:
		fmt.Println("What the hell is this value?")
	}
}