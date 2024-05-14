package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition for two positive numbers", 2, 2, "+", 4, ""},
		{"addition for two negative numbers", -2, -2, "+", -4, ""},
		{"addition with zero", 2, 0, "+", 2, ""},
		{"addition of two zero", 0, 0, "+", 0, ""},
		{"addition of a positive and a negative number", 2, -2, "+", 0, ""},
		{"subtraction of two positive numbers", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"multiplication", 2, 3, "*", 6, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, `division by zero`},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			assertions := assert.New(t)

			assertions.Equal(result, d.expected, "They are equal")
			//if result != d.expected {
			//	t.Errorf("Expected %d, got %d", d.expected, result)
			//}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`",
					d.errMsg, errMsg)
			}
		})
	}

}
