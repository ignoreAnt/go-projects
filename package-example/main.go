package main

import (
	"fmt"
	nfmt "go-projects/package-example/do-format"
	"go-projects/package-example/math"
	"go-projects/package-example/name"
	"go-projects/package-example/util"
)

func main() {

	num := math.Double(2)
	output := nfmt.Number(num)
	fmt.Println(output)

	util.ExtractNames("hello")
	util.FormatNames("hello")

	name.Extract("hello")
	name.Format("Hello")

}
