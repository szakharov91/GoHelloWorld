package main

import "fmt"

func main() {
	println("Hello World")

	println("I started to learn GO lang")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}

	fmt.Println(s)

	f := []int{2, 3, 5, 7, 11, 13}

	f = f[1:4]

	printSlice(f)

	var t []int
	fmt.Println(t, len(t), cap(t))

	if t == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
