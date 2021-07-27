package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20

	p := fmt.Println

	p(x == y)
	p(x != y)
	p(x < y)
	p(x <= y)
	p(x > y)
	p(x >= y)
}