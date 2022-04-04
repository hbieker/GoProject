package main

import (
	"fmt"
)

//Uppercase variable in package is globally visible
//Lowercase variable is scoped to this package only

//Used for shadowing example below
var d float32 = 40

//var block example
var (
	firstName string = "John"
	lastName  string = "Doe"
	age       int    = 32
)

func main() {

	var a int
	a = 10
	fmt.Println(a)

	var b int = 20
	fmt.Println(b)

	c := 30 // := Cannot be used at package level
	fmt.Println(c)
	fmt.Printf("Value and Type is %v, %T\n", c, c)

	//Shadow variable is initially the package level value
	fmt.Printf("Value and Type is %v, %T\n", d, d)
	d := 41
	fmt.Printf("Value and Type is %v, %T", d, d)

}
