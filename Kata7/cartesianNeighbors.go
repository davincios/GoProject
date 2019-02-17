package main

import "fmt"

/*
A Cartesian coordinate system is a coordinate system that specifies each point uniquely in a plane by a pair of numerical coordinates,
which are the signed distances to the point from two fixed perpendicular directed lines, measured in the same unit of length.

The сoordinates of a point in the grid are written as (x,y).
Each point in a coordinate system has eight neighboring points. Provided that the grid step = 1.

It is necessary to write a function that takes a coordinate on the x-axis and y-axis and returns a list of all the neighboring points.
Points inside list don't have to been sorted (any order is valid).

For Example:

CartesianNeighbor(2,2) -> {{1,1},{1,2},{1,3},{2,1},{2,3},{3,1},{3,2},{3,3}}
CartesianNeighbor(5,7) -> {{6,7},{6,6},{6,8},{4,7},{4,6},{4,8},{5,6},{5,8}}
*/

func CartesianNeighbor(x, y int) [][]int {
	// Create a dynamically sized array
	// Each single combination will have two values in its array, and those combinations will be appended to the answer array.
	// Therefore we need to make use of a nested integer array.
	neighbourArray := make([][]int, 0, 8)

	// We start looping x, and start at the lower neighbour and proceed to the current and upper neighbour.
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			// Skip if both values are equal to the initial values.
			if i == x && j == y {
				// A continue statement in Go skips the current iteration, instead of terminating.
				continue
			}

			// Define the current coordinate combination
			combination := []int{i, j}

			// Append the combination to the list with all combinations
			neighbourArray = append(neighbourArray, combination)
		}

	}
	return neighbourArray
}

func main() {
	fmt.Println(CartesianNeighbor(2, 2)) // -> {{1,1},{1,2},{1,3},{2,1},{2,3},{3,1},{3,2},{3,3}}
	fmt.Println(CartesianNeighbor(5, 7)) // -> {{6,7},{6,6},{6,8},{4,7},{4,6},{4,8},{5,6},{5,8}}
}
