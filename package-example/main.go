package main

import (
	"fmt"
	format "go-projects/package-example/do-format"
	"go-projects/package-example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
