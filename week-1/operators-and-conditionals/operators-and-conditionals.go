package main

import (
	"fmt"
	"math"
)

// if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if with short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if and else
func isOdd(n int8) bool {
	if v := n % 2; v == 0 {
		return true
	} else {
		return false
	}
}

// switch
func carCompany(carName string) string {
	switch carName {
	case "kwid":
		return "Renautl"
	case "mobi":
		return "Fiat"
	default:
		return "Other"
	}
}

func main() {
	fmt.Println(sqrt(25), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(isOdd(3), isOdd((4)))
	fmt.Println(carCompany("kwid"), carCompany("mobi"), carCompany("hb20"))

}
