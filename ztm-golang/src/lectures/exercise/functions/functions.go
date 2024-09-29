//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetPerson(name string) {
	fmt.Println("Hi", name, "!")
}

func hiThere() string {
	return "Hi there!"
}

func addThreeNumbers(number1, number2, number3 int) int {
	return number1 + number2 + number3
}

func returnOne() int {
	return 1
}

func returnTwoAndFour() (int, int) {
	return 2, 4
}

func main() {
	greetPerson("Andre")
	fmt.Println(hiThere())

	two, four := returnTwoAndFour()

	totalNumbers := addThreeNumbers(returnOne(), two, four)
	fmt.Println(totalNumbers)

}
