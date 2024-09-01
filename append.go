package main

import "fmt"

func main() {
	var s []int
	printSliceAppend(s)

	s = append(s, 0)
	printSliceAppend(s)

	s = append(s, 1)
	printSliceAppend(s)

	s = append(s, 2, 3, 4)
	printSliceAppend(s)
}

func printSliceAppend(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
