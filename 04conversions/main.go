package main

/* Convert string to int
* strconv.Atoi() function is used to convert a string to an integer
* The function returns two values, the converted integer and an error
* If the conversion fails, the error will be non-nil
* If the conversion is successful, the error will be nil
* The error should be checked before using the converted integer
* The error can be checked using the if statement
* If the error is nil, the conversion was successful
* If the error is non-nil, the conversion failed
* The converted integer can be used after the error check
* The converted integer can be stored in a variable
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to the program!")
	fmt.Println("Please enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	
	input, _ := reader.ReadString('\n')
	fmt.Print("Thank you for entering: ", input, "\n")

	// number, err := strconv.Atoi(strings.TrimSpace(input))	// To convert a string to an integer
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err == nil {
		fmt.Println("Conversion successful!")
		fmt.Println("Added 10 to your given number: ", number+10)
	} else {
		fmt.Println("Conversion failed!")
		fmt.Println("The error is: ", err)
	}
}