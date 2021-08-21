// Declare package
package main

//import the extension cooks
import (
	"fmt"
)

// main function call
func main() {
	fmt.Println("I am revolving - ")

	var enum int8
	var combinedName string

	for i := 0; i < 129; i++ {
		enum += 1
		fmt.Println(enum)
	}

	//Call function
	myfun()

	//call function
	myfun1("ansuman", "padhy")
	fmt.Println(enum)

	//call function
	combinedName = myfun2("IBM", "INDIA")
	fmt.Println(combinedName)

	//call function
	combinedName, err := myfun3("IBM", "INDIA")
	fmt.Println(combinedName)
	if err != nil {
		fmt.Println(combinedName)
	}
	fmt.Println("I don't have any value from myfun3")
}

//No Argument and No returns function call
func myfun() {
	println("I am myfun block")

	var apple, orange, fruit int
	apple = 5
	orange = 5
	fruit = apple + orange
	fmt.Println(fruit)
}

//Passing arguments and No returns function call

func myfun1(a string, b string) {
	var fullname string
	fullname = a + " " + b
	println(fullname)
}

//Passing argument and returns function call
func myfun2(a string, b string) string {
	var fullname string
	fullname = a + " " + b
	println("Am I retruning?")
	return fullname
}

//Passing argument and returns function call
func myfun3(a string, b string) (string, error) {
	var fullname string
	fullname = a + " " + b
	println("Am I retruning?")
	return fullname, nil
}
