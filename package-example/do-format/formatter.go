package format

import "fmt"

func Number(num int) string {
	return fmt.Sprintf("The number is %d", num)
}

func init() {
	fmt.Println("init function called")
}
