package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("--- File Handling in Golang ---")
	
	content := "Hello, World! \n This is a file handling in Golang. We are writing this content to a file."
	
	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}
	
	
	length, err := io.WriteString(file, content)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Length of the file is: ", length)
	defer file.Close()
	readFile("./test.txt")
}

func readFile(filename string) {
	// Shown in the video but it is now depricated, look for another way to read file.
	// inDataByteFormat, err := ioutil.ReadFile(filename)
	file, err := os.Open(filename)
	data, err := io.ReadAll(file)
	CheckErrorNill(err)
	// fmt.Println("Data read from file(raw): \n", inDataByteFormat)
	// fmt.Println("Data read from file(in string format): \n", string(inDataByteFormat))
	fmt.Println("Data read from file(raw): \n", data)
	fmt.Println("Data read from file(in string format): \n", string(data))
}

func CheckErrorNill (err error){
	if err != nil {
		panic(err)
	}
}