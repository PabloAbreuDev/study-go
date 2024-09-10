package main

import (
	"fmt"
)

// Passing parameter as reference
func changeValue(p *int){
	*p = 20
}

// Pointer with structs
type Person struct {
	Name string
	Age int
}

func changeAge(p *Person, newAge int){
	p.Age = newAge
}

func main(){
	fmt.Println("Pointers")

	x := 10
	p := &x

	fmt.Println(x)
	fmt.Println(*p)

	y := 15
	fmt.Println(y)
	changeValue(&y)
	fmt.Println(y)


	person := Person{Name: "Alice", Age: 30}
	changeAge(&person, 12)
	fmt.Println(person)
	
}
