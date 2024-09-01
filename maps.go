package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": Vertex{},
}

func main() {
	fmt.Println(m["Bell Labs"])

	fmt.Println(m["Google"])

	fmt.Println(m)
}
