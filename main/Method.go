package main

import (
	"fmt"
	"math"
)

// method is not function in Go

type Vertex struct {
	X, Y float64
}

// This is a method of struct Vertex
// The "(v Vertex)" is the receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This the function with the same functionality of above method
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Non-struct type could also have methods
type MyFloat float64

func (mf MyFloat) demo() float64 {
	return float64(mf)
}

// pointer receiver, returns the pointer to the object that are used usually for modifying the underlying value or avoid copying
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

/**
If here is receiver argument, it will avoid copy the v.X and v.Y in memory
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(3.0)
	fmt.Println(f.demo())

	// if the above Scale method's receiver is "v Vertex", the result gonna be different
	// although the "v" is value instead of a pointer, the method with the pointer receiver is called automatically, that is: v.Scale(5) as (&v).Scale(5) for the convenience.
	v.Scale(10)
	fmt.Println(v.Abs())
}
