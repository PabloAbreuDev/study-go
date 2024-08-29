package main

import (
	"fmt"
)

func modifyMap(m map[string]int) {
	m["Alice"] = 10
}

func main() {
	// Maps are similar to objects in javascript
	var m map[string]int

	fmt.Println("A nil map: ", m)

	// Initialized map
	m2 := make(map[string]int)

	fmt.Println("Create and initialize an empty map: ", m2)

	// Initialize a map with values
	m3 := map[string]int{
		"Alice": 23,
		"Bob":   34,
	}

	fmt.Println(m3)

	// Inserting data
	m4 := make(map[string]int)

	m4["John"] = 18

	m4["Leon"] = 37

	fmt.Println(m4["Leon"])

	// Removing elements in a map
	delete(m4, "Leon")

	fmt.Println(m4)

	// Iterating a map
	for key, value := range m3 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	m5 := map[string]int{
		"Alice": 25,
	}

	fmt.Println(m5)

	modifyMap(m5)

	fmt.Println(m5)

}
