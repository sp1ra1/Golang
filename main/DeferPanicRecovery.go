package main

import "fmt"

// A deferred function's arguments are evaluated when the defer statement is evaluated.
// "1"
func deferRule1() {
	i := 1
	defer fmt.Println(i)
	i++
	return
}

// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func deferRule2() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
}

// Deferred functions may read and assign to the returning function's named return values.
// Deferred function increments the return value i after the surrounding function returns. Thus, returns 2
func deferRule3() (i int) {
	defer func() { i++ }()
	return 1
}


func main() {
	deferRule1()
	deferRule2()
	fmt.Println(deferRule3())

	// https://blog.golang.org/defer-panic-and-recover
	// Panic and Recover example
}
