package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z string
	A bool
}

func main() {
	fmt.Println(Vertex{1, 2, "test", false})
}
