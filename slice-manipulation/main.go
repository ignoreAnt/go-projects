// Program to show the manipulation of the string slice
package main

import "fmt"

// Slice Manipulation
// Objective: Create a function that receives a slice of strings and a string to remove.
//
//	The function should return a new slice with all occurrences of the string removed.
//
// Concepts: Slices, iteration, functions.
// Challenge: Ensure your function is efficient and does not unnecessarily copy elements.
func main() {
	strs := []string{"boy", "class", "board", "boy", "class", "girl"}
	fmt.Println(removeAllOccurrences(strs, "boy"))
	fmt.Println(removeAllOccurrences(strs, "girl"))
	fmt.Println(removeAllOccurrences(strs, "duster"))
	fmt.Println(removeAllOccurrencesInPlace(strs, "boy"))
	fmt.Println(strs)
}

// removeAllOccurrences removes all the occurrences of elem from values
// Creates a new slice and thus original slice is not changed
// If the elem supplied is not present in the slice, nothing is removed
func removeAllOccurrences(values []string, elem string) []string {
	var output []string
	for _, v := range values {
		if v != elem {
			output = append(output, v)
		}
	}

	return output
}

// Challenge: Ensure your function is efficient and does not unnecessarily copy elements.
// Removes the elements in place without needing an extra slice space. Changes the original slice
func removeAllOccurrencesInPlace(s []string, toRemove string) []string {
	i := 0
	for _, elem := range s {
		if elem != toRemove {
			s[i] = elem
			i++
		}
	}

	return s[:i]
}
