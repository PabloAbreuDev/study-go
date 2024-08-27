package main

import (
	"fmt"
)

func modifyArray(a *[3]int){
	a[0] = 100
}

func main() {
	fmt.Println("Foo")

	var arr [5]int

	fmt.Println(arr)

	var arrTwo = [5]int{1, 2, 3, 4, 5} // Initialize an array with 5 positions

	fmt.Println(arrTwo)

	arrThree := [...]int{1,2,3} // Go defines the array length

	fmt.Println(arrThree)

	// Accessing elements
	arrFour := [3]int{10,20,30}

	fmt.Println(arrFour[0])

	// Array length
	fmt.Println(len(arrFour))

    // Iterate an array with for
	for i := 0; i < len(arrFour); i++ {
		fmt.Println("With for:",arrFour[i])
	}

	// Iterate an array with range
	for index,value := range arrFour {
		fmt.Printf("Index %d, Value %d\n", index, value)
	}

	// Matrix
	var matrix [2][3]int

	matrix = [2][3]int {
		{1,2,3},
		{4,5,6},
	}

	fmt.Println(matrix[1][2])


	arrForFunction := [3]int{1,2,3}
	modifyArray(&arrForFunction)
	fmt.Println(arrForFunction)

	/////--------------Slices-----------/////

	var s []int // Declaring a slice
	fmt.Println(s)

	s2 := []int{1,2,3}
	fmt.Println(s2)
	
	arr5 := [5]int{1,2,3,4,5}
	s5 := arr5[1:4]
	fmt.Println(s5)

	s5 = append(s5,10,12,20 )
	fmt.Println(s5)



}
