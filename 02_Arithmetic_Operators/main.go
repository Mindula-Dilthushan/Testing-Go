//Alpha X Software Company
//Mindula Dilthushan
//Testing-Go
//module-02
//21-07-23
package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20
	var z int

	z = x + y
	fmt.Printf("x + y = %d\n", z)

	z = x - y
	fmt.Printf("x - y = %d\n", z)

	z = x * y
	fmt.Printf("x * y = %d\n", z)

	z = x / y
	fmt.Printf("x / y = %d\n", z)

	z = x % y
	fmt.Printf("x %% y = %d\n", z)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)

}