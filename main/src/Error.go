package main

import (
	"fmt"
	"time"
)

// The "error" type is a built-in interface in "fmt" like Stringer
//type error interface {
//	Error() string
//}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"id didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
