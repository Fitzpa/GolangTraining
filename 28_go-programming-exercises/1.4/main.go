// Hands-on exercise #4
// FYI - nice documentation and new terminology “underlying type”
// https://golang.org/ref/spec#Types
// For this exercise
// Create your own type. Have the underlying type be an int.
// create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
// in func main
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign 42 to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”
package main

import "fmt"

type cat int

var x cat

func main() {
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)
	x = 42
	fmt.Println(x)
}
