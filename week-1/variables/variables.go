package main

import "fmt"

var c, python, java bool

func main(){
	// Variables without initializers
	var i int
	fmt.Println(i, c, python, java)

	// Variables with initializers
	var x, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(x,j,c,python,java)

	// Short variables declartion
	var n, z = 1, 2
	l := 3
	fmt.Println(n,z,l)
	
}