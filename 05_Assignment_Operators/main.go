package main

import "fmt"

func main() {

	var x, y = 10, 20
	x = y
	fmt.Println(x)

	x = 10
	x += y
	fmt.Println(x)

	x = 10
	x -= y
	fmt.Println(x)

	x = 10
	x *= y
	fmt.Println(x)

	x = 20
	x /= y
	fmt.Println(x)

	x = 20
	y = 3
	x %= y
	fmt.Println(x)
}