package main

import (
	"fmt"
	"time"
)

/*

* Time is a package in Go that provides functionality for measuring and displaying time.
* Time is measured in nanoseconds.
* Time is a struct in Go.
* Time has a method called 'Now()' which returns the current time.
* Time has a method called 'Date()' which returns a time corresponding to the given year, month, day, hour, minute, second, and nanosecond.
* Time has a method called 'Format()' which formats the time according to the layout.
* When using Format() method, the layout must use the reference time 'Mon Jan 2 15:04:05 MST 2006' to show the pattern with which to format/parse a given time/string.
* Time has a method called 'Add()' which returns the time corresponding to the given duration added to the time.

 */

func main(){
	fmt.Println("Objective: To learn about time in Go")
	
	presentTime := time.Now()
	fmt.Println("Present time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

// Create a Time
	createdTime := time.Date(2000, time.February, 10, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created time is: ", createdTime)
	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Mon"))
}