package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

type Input struct {
	A string
	B string
}

type some interface {
	string() string
}

type COut struct {
	frequencyCount map[rune]int
}

func (o COut) string() string {
	return ""
}

func (o COut) process() error {
	return nil
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	ab := newABProcessor()
	ab.start(ctx, data)
	inputC, err := ab.wait(ctx)
	if err != nil {
		return COut{}, err
	}

	// TODO : refactor it for better tests
	c := newCProcessor()
	c.start(ctx, inputC)
	out, err := c.wait(ctx)
	return out, err
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("expected input to be processed")
		os.Exit(1)
	}
	cout, err := GatherAndProcess(context.Background(), Input{
		A: os.Args[1],
		B: os.Args[2],
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cout)
}
