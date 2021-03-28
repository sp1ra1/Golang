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
	p4 := &Pair{1, 2}
	fmt.Println(p1, p2, p3, *p4)
}

func arrays() {
	var a [2]int
	a[0] = 0
	a[1] = 1
	fmt.Println(a)
	primes := [6]int{1, 2, 3, 5, 7, 11}
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
	var slice = []int{4, 5}
	s := []struct {
		row int
		col int
	}{
		{1, 2},
		{3, 4},
	}
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(s)
}

// todo
func sliceLengthAndCap() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)

	s = s[:0]
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)

	s = s[:4]
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)

	s = s[2:]
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)

	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("gg!")
	}
}

// make(type, len, cap)
func makeSlice() {
	a := make([]int, 5)
	fmt.Println(a)

	b := make([]bool, 5, 10)
	fmt.Println(len(b), cap(b))
}

// append(s []T, ...T) []T
func appendSlice() {
	var s []string
	s = append(s, "gg")
	s = append(s, "my", "friend")
	fmt.Println(s)
}

func rangeStatement() {
	a := []int{1, 2, 3, 4, 5, 6}
	for i, v := range a {
		fmt.Println(i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}
}

func mapDemo() {
	var a map[string]bool
	a = make(map[string]bool)
	a["gg"] = true
	fmt.Println(a["gg"])

	var m = map[string]float32{
		"gg":     3.0,
		"friend": 4.0,
	}
	fmt.Println(m)
}

func mapOperation() {
	m := make(map[string]int)
	m["answer"] = 42
	delete(m, "answer")
	// if the key is in the map, return the corresponding value and true, else return zero value of the type and false
	v, ok := m["answer"]
	fmt.Println(v, ok)
}

func funcParameter(x, y float64) float64 {
	return x + y
}

func wrapper(f func(float64, float64) float64) float64 {
	return f(3.0, 4.0)
}

// function closure
// each closure is bound to its own "sum" variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
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
	appendSlice()

	rangeStatement()

	mapDemo()
	mapOperation()

	fmt.Println(wrapper(funcParameter))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
