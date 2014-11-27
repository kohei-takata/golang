package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z string
	A bool
}

func main() {
	v := Vertex{1, 2, "test", false}
	v.X = 4
	fmt.Println(v.A)
}
