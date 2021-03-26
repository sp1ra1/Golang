package main

import "fmt"

func pointer() {
	var ptr *int
	i := 12
	ptr = &i
	fmt.Println(*ptr)
	*ptr = 20
	fmt.Println(*ptr)
}

type Pair struct {
	row int
	col int
}

func structInit() {
	p1 := Pair{1, 2}
	p2 := Pair{row: 1, col: 2}
	p3 := Pair{}
	p4 := &Pair{1,2}
	fmt.Println(p1, p2, p3, p4)
}

func arrays()  {
	var a [2]int
	a[0] = 0
	a[1] = 1
	fmt.Println(a)
	primes := [6]int{1,2,3,5,7,11}
	fmt.Println(primes)
}

// A slice does not store any data, it just describes a section of an underlying array.
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
// Other slices that share the same underlying array will see those changes.
func slices() {
	names := [4]string{
		"Paul",
		"Jon",
		"Charity",
		"JT",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)
	b[0] = "Cyrus"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceVerusArray() {
	// This is an array
	var array = [3]int{1, 2, 3}
	// This is a slice (describe a section of underlying data)
	var slice = []int{4,5}
	s := []struct{
		row int
		col int
	} {
		{1, 2},
		{3, 4},
	}
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(s)
}

func sliceLengthAndCap() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s);

	s = s[:0]
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s);

	s = s[:4]
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s);

	s = s[2:]
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s);

	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("gg!")
	}
}

func makeSlice() {
	a := make([]int, 5)
	fmt.Println(a)
}

func main() {
	pointer()
	fmt.Println(Pair{1, 2})
	p := Pair{3, 4}
	fmt.Println(p.col)

	// struct pointer( *(ptr).col = ptr.col )
	pair := Pair{10, 20}
	ptr := &pair
	fmt.Println(ptr.row)

	structInit()
	arrays()
	slices()
	sliceVerusArray()
	sliceLengthAndCap()
	makeSlice()
}