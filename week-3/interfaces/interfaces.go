package main

import (
	"fmt"
)

// Interface for Speaker with method "Speak"
type Speaker interface {
	Speak() string
}
// Struct Person that implements Speaker interface
type Person struct {
	Name string
}

// Speak method from struct person
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}


type Animal interface {
    Speak() string
}

type Dog struct {}
type Cat struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func makeSound(a Animal){
	fmt.Println(a.Speak())
}

func main(){
	var s Speaker
	p := Person{Name: "Alice"}

	s = p

	fmt.Println(s.Speak())

	// In go, is not necessary to declare an implementation of a interface.
	d := Dog{}
	c := Cat{}

	fmt.Println(d.Speak())
	fmt.Println(c.Speak())

	makeSound(d)
	


}



