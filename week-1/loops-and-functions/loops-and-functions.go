package main

import (
	"fmt"
)

func printAndSum(arr [5]int) {
	sum := 0
	for _, value := range arr {
		fmt.Println("Number: ", value)
		sum += value
	}
	fmt.Println(sum)
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	printAndSum(numbers)
}
