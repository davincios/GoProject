package main

import "fmt"

func PrinterError(s string) string {
	// First convert string into runes
	runes := []rune(s)
	// Create error container to store the errors in
	errors := []string{}
	// Loop through the string to check for values higher than m=109
	for _, rune := range runes {
		if rune > 109 {
			x := string(rune)
			errors = append(errors, x)
		}
	}

	// Create string to return the amount of errors divided by total mount.
	answer := fmt.Sprintf("%d/%d", len(errors), len(string(runes)))
	return answer
}

func main() {
	fmt.Println(PrinterError("abcdemxyz"))                                                // should return "3/9"
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")) // should return "3/56"

}
