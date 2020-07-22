package main

import (
	"fmt"
	"strconv"
)

//global vatriable exposed to other package to
var I int = 56

//global variable but limited to package only
var g int = 49
var (
	name string = "Tarun"
	age int = 25
)

func main() {
	var i int= 42
	var j1 string = string(i)
	
	fmt.Println(j1)
	
	j1 = strconv.Itoa(i)
	fmt.Println(j1)

	a := 55.
	var j float32 = 32.
	j = float32(i)
	fmt.Println(j)
	fmt.Println(i)
	fmt.Println(g)
	fmt.Printf("%v, %T", a, a)
}
