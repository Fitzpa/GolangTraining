package main

import "fmt"

func main() {
	a := 42
	// print the value of 'a' in base 10, binary and hexidecimal
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}
