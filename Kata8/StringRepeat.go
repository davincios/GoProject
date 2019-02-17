package main

import (
	"fmt"
	"strings"
)

func repeatStr(repetitions int, value string) string {
	var answer []string
	for i := 0; i < repetitions; i++ {
		answer = append(answer, value)
	}
	fmt.Println(strings.Join(answer, ""))
	return ""
}

func repeatIt(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

// Alternative approach with Go's built in repeat function

func main() {
	repeatStr(6, "I")     // "IIIIII"
	repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"

	fmt.Println("Alternative approach with Go's built in repeat function")
	fmt.Println(repeatIt(12, "Hello"))

}
