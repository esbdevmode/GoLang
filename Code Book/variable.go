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

	for i := 0; i < 129; i++ {
		enum += 1
		//fmt.Println(enum)
	}

	//Call function
	myfun()

	//call function
	myfun1("ansuman", "padhy")
	fmt.Println(enum)
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
