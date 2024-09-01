package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSliceModify("a", a)
}

func printSliceModify(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
