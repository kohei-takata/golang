package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("welcome")
	fmt.Println("The time is", time.Now())
	fmt.Println(os.Open("doc.go"))
	fmt.Println(net.Dial("tcp", "google.com"))
}
