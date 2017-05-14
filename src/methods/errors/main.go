package main

import (
	"fmt"
	"time"
	"math"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	sqrt, sqrt_err := Sqrt(2)
	if sqrt_err != nil {
		fmt.Println(sqrt_err)
	} else {
		fmt.Println(sqrt)
	}

	sqrt, sqrt_err = Sqrt(-2)
	if sqrt_err != nil {
		fmt.Println(sqrt_err)
	} else {
		fmt.Println(sqrt)
	}
}
