package main

import (
	"fmt"
)

func century(year int) int {
	// Finish this :)
	century := year / 100

	// Check if the century is correct
	if year-(century*100) > 0 {
		century = century + 1
		fmt.Println("The year is", year, "The century is", century)
		return century
	} else if (century*100)-year == 0 {
		fmt.Println("The century is one year less")
		fmt.Println("The year is", year, "The century is", century)
		return century
	} else {
		fmt.Println("Something strange is going on ---- Errror")
		return century
	}
}

func main() {
	century(1705) // Has to return 18
	century(1601) // Has to return 17
	century(1900) // has to return 19
	century(2000) // Has to return 20
}
