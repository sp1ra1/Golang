package main

import (
	"fmt"
	"math"
)

// The declared interface
type Abser interface {
	Abs() float64
}

type MyFloatAlias float64

type VertexAlias struct {
	X float64
	Y float64
}

// A type implement an interface by implementing its method without any "implement" keywords like Java
func (f MyFloatAlias) Abs() float64 {
	return float64(f)
}

func (v *VertexAlias) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// String() method is declared in interface called "Stringer" in fmt package
// The type that implement this method would describe itself as a string
// It is a little bit like the toString() method in Object class in Java
func (v VertexAlias) String() string {
	return fmt.Sprintf("%v (%v years)", v.X, v.Y)
}

func main() {
	v := VertexAlias{3, 4}
	fmt.Println(v.Abs())

	mf := MyFloatAlias(3)
	fmt.Println(mf.Abs())

	// empty interface could hold values of any type. It is a little bit like Object in Java
	var i interface{}
	i = 42
	fmt.Println(i)
	i = "gg"
	fmt.Println(i)
	//For example, fmt.println() takes any number of arguments of type interface{}
	/**
	func Println(a ...interface{}) (n int, err error) {
		return Fprintln(os.Stdout, a...)
	}
	*/

	// t := i.(T)
	// This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	g, ok := i.(byte)
	fmt.Println(g, ok)

	// If i does not hold a T, the statement will trigger a panic.
	//n := i.(byte)
	//fmt.Println(n)

	fmt.Println(v)
}
