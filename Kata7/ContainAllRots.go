package main

import (
	"fmt"
)

func ContainAllRots(s string, arr []string) bool {
	// accoutn for empty string case immediately
	if s == "" {
		return true
	}

	// create a dynamically sized container for all the rotations generated for the passed string
	generated := make([]string, 0)

	// also create a container for all the valid rotations found in the passed array
	rotationsList := make([]string, 0)

	// first, convert the string in into a slice of runes
	// Rune literals are just 32-bit integer values. They are untyped so their type can chanfe
	runes := []rune(s)

	// then, iterate over that slice to generate complete list of rotations for the passed string
	for _, char := range runes {
		// pop the first element off the slice ...
		char, runes = runes[0], runes[1:]

		// append the element to the end of the slice
		runes = append(runes, char)

		// turn the slice of runes back into a string ...
		rotation := string(runes)

		// ... and append the rotation to the list of known rotations
		generated = append(generated, rotation)

		// next, take that roation and check to see if it exists in the passed array argument
		for _, rot := range arr {
			// if it is in there, append it to the slice of rotations found in the passed array
			if rotation == rot {
				rotationsList = append(rotationsList, rotation)
			}
		}
	}

	/* The array lengts of the generated rotations and the provided rotations will be
	Equal if the passed array does contain all valid rotations of the passed string
	*/

	if len(generated) == len(rotationsList) {
		fmt.Println("The provided array list contains all know rotations")
		return true
	} else {
		fmt.Println("The provided array list DOES NOT contain all know rotations")
		return false
	}

}

func main() {
	ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"})                                                       // Should return true
	ContainAllRots("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}) // Should return false
}
