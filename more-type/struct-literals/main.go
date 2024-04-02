package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{Y: 2}
	v4 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, v4, p)
}
