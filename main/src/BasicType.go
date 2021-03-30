package main

import (
	"fmt"
)

func holyShit(x int, y int) int {
	return x + y
}

// omit the type from all but the last when two or more consecutive named function parameters
func holyShitAnotherVersion(x, y int) int {
	return x + y
}

// multiple return results
func swap(x, y string) (string, string) {
	return y, x
}

// Return statement without arguments returns the named return values
func nakedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	fmt.Println(holyShit(1, 2) == holyShitAnotherVersion(2, 1))
	fmt.Println(swap("gg", "my friend"))
	fmt.Println(nakedReturn(1))
	var golang int
	fmt.Println(c, python, java, golang)

	// if has initializer, the type could be omitted
	var cpp, js, py = true, 0, "gg"
	fmt.Println(cpp, js, py)

	// inside a function, var equals to :=
	perl, lua := "jesus", true
	fmt.Println(perl, lua)

	// zero value of a string is ""
	var str string
	fmt.Println(str)

	// When the right hand side contains an untyped numeric constant, the new variable's type could vary
	var i = 3
	fmt.Printf("i is of type %T\n", i)
	var j = 3.00
	fmt.Printf("j is of type %T\n", j)

	// Constants cannot be declared using the := syntax.
	const linus = "yyds"
	fmt.Println(linus)
}
