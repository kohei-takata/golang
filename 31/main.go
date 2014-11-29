package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	i := 3
	j := 4
	fmt.Println("p ==", p)
	fmt.Printf("p[1:%d] == %d\n", j, p[1:j])
	fmt.Printf("p[:%d] == %d\n", i, p[:i])
	fmt.Printf("p[%d:] == %d\n", j, p[j:])
}
