package main

import "fmt"

func EvenOrOdd(number int) string {
	isEven := number%2 == 0
	if isEven {
		fmt.Println(number, "Even")
		return "Even"
	} else {
		fmt.Println(number, "Odd")
		return "Odd"
	}
}

func main() {
	EvenOrOdd(1705) // Has to return odd
	EvenOrOdd(1601) // Has to return odd
	EvenOrOdd(1900) // has to return even
	EvenOrOdd(2000) // Has to return even
}
