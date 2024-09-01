package main

import "fmt"

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return second - first
	}
}
