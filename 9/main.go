package main

import "fmt"

func swap(x, y int) (a, b, c, d int) {
	a = x / y
	b = x * y
	c = x + y
	d = x - y
	return
}

func main() {
	a, b, c, d := swap(4, 2)
	fmt.Println(a, b, c, d)
}
