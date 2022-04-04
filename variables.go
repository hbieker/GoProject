package main

import (
	"fmt"
	"strconv"
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
	fmt.Printf("Value and Type is %v, %T\n", d, d)

	//convert from one variable type to another
	var e int = 50
	fmt.Printf("Value and Type is %v, %T before conversion \n", e, e)

	var f float32
	f = float32(e)
	fmt.Printf("Value and Type is %v, %T after conversion\n", f, f)

	//Convert int to String. ex: for logging. Otherwise it converts to the ascii char
	var g int = 60
	fmt.Printf("Value and Type is %v, %T before conversion to string\n", g, g)

	var h string
	h = strconv.Itoa(g)
	fmt.Printf("Value and Type is %v, %T after conversion\n", h, h)

}
