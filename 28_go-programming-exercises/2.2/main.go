package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	// print the value of 'a' in base 10, binary and hexidecimal
	fmt.Printf("%v\t", a)
	fmt.Printf("%v\t", b)
	fmt.Printf("%v\t", c)
	fmt.Printf("%v\t", d)
	fmt.Printf("%v\t", e)
	fmt.Printf("%v\t", f)
}
