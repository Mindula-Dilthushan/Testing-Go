package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20
	var z int = 0

	z = x & y
	fmt.Println(z)

	z = x | y
	fmt.Println(z)

	z = x ^ y
	fmt.Println(z)

	z = x << 1
	fmt.Println(z)

	z = x >> 1
	fmt.Println(z)
}