package main

import (
	"fmt"
	"time"
)

func numbersFunc() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func lettersFunc() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c", l)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	// This both functions runs concurrently when we use "go" label
	go numbersFunc()
	go lettersFunc()
	time.Sleep(time.Second * 5)
	fmt.Println("End of program")
}
