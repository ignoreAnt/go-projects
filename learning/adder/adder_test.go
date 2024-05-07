package adder

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond also uses stuff set up in TestMain", testTime)
}

func Test_addNumber(t *testing.T) {
	// Arrange
	x := 2
	y := 3
	expected := 5

	// Act
	result := addNumber(x, y)

	// Assert
	if expected != result {
		t.Fatalf("incorrect output expected %d, but got %d: ", expected, result)
	}
}

func Test_addZero(t *testing.T) {
	// Arrange
	x := 2
	y := 0
	expected := 2

	// Act
	result := addNumber(x, y)

	// Assert
	if expected != result {
		t.Fatalf("incorrect output expected %d, but got %d: ", expected, result)
	}
}

func Test_addNegative(t *testing.T) {
	// Arrange
	x := 2
	y := -5
	expected := -3

	// Act
	result := addNumber(x, y)

	// Assert
	if expected != result {
		t.Fatalf("incorrect output expected %d, but got %d: ", expected, result)
	}
}

func Test_addNegativeV2(t *testing.T) {
	// Arrange
	x := 2
	y := -5
	expected := -3

	// Act
	result := addNumber(x, y)
	t.Fail()

	// Assert
	if expected != result {
		t.Fatalf("incorrect output expected %d, but got %d: ", expected, result)
	}
}
