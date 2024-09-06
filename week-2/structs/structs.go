package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{"Alice", 25}
	fmt.Println(p)

	p2 := Person{Name: "Pablo", Age: 28}
	fmt.Println(p2)

	p3 := Person{Name: "John", Age: 18}
	fmt.Println(p3)
	p3.Age = 35
	fmt.Println(p3)

	type Address struct {
		City  string
		State string
	}

	type Employee struct {
		Name string
		Age  int
		Address
	}

	e := Employee{Name: "John", Age: 35, Address: Address{City: "New York", State: "NY"}}

	fmt.Println(e)

	type Company struct {
		Name    string
		Founder Person
	}

	c := Company{Name: "TechCorp", Founder: Person{Name: "Alice", Age: 25}}
	fmt.Println(c.Founder.Name)

}
