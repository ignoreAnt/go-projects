// Packge main runs the main package
// This is for the array transformation exercise
package main

import (
	"fmt"
)

// Array Transformation
// Objective: Write a function that takes a slice of integers and returns a slice containing only the even numbers from the slice.
// Concepts: slices, functions.
// Challenge: Try to do it in-place (modifying the original slice) and then try returning a new slice.
func main() {
	nums := []int{4, 57, 2, 7, 3, 25, 36, 100, 87, 29}
	fmt.Println(filterEvenNumber(nums))
	fmt.Println(filterEvenInPlace(nums))
}

// arrayTrans function takes a slice an input and return only the number which is even
// vals in considered un sorted
func filterEvenNumber(vals []int) []int {
	var filtered []int

	for i := 0; i < len(vals); i++ {
		if vals[i]%2 == 0 {
			filtered = append(filtered, vals[i])
		}
	}

	return filtered
}

// filterEvenInPlace function takes a slice as input and save the number in place
func filterEvenInPlace(vals []int) []int {
	i := 0
	for j := 0; j < len(vals); j++ {
		if vals[j]%2 == 0 {
			vals[i] = vals[j]
			i++
		}
	}

	// Fill the remaining elements with zero, not the best solution
	for ; i < len(vals); i++ {
		vals[i] = 0
	}
	
	// better return a slice
	return vals[:i]
}
