package main

import (
	"fmt"
)

/*
CONDITIONS
Float parameter "h" in meters must be greater than 0
Float parameter "bounce" must be greater than 0 and less than 1
Float parameter "window" must be less than h.
*/

func BouncingBall(h, bounce, window float64) int {
	// Check if initial conditions are adhered to.
	if h < 0 || bounce <= 0 || 1 <= bounce || h < window {
		return -1
	}
	var count int = -1

	for ; h > window; h *= bounce {
		count += 2
	}

	return count
}

func main() {
	fmt.Println(BouncingBall(40, 0.4, 10)) //  3
	fmt.Println(BouncingBall(10, 0.6, 10)) // -1
	fmt.Println(BouncingBall(40, 1, 10))   // -1
	fmt.Println(BouncingBall(5, -1, 1.5))  // -1
}
