package main


import (
	"fmt"
)

func applyOperation(a, b int, operation func(int,int) int) int {
	return operation(a,b)
}

func main(){
	fmt.Println("First class functions")
	// Defining a function to sum and other one to sub
	add := func(x,y int) int {
		return x + y
	}

	sub := func (x,y int) int {
		return x - y
	}

	// Passing functions as parameters
	result1 := applyOperation(5,3,add)
	result2 := applyOperation(3,5,sub)

	fmt.Println(result1)
	fmt.Println(result2)
}
