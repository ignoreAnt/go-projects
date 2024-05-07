package math

import "fmt"

func init() {
	fmt.Println("init called from double")
}

func init() {
	fmt.Println("second init")
}

func init() {
	fmt.Println("third init")
}

func Double(a int) int {
	return a * 2
}
