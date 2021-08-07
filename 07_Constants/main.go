package main

import (
	"fmt"
	"math"
)

const PI float64 = 3.14159265359

func main() {
	fmt.Println(PI)

	//A const statement can be written in anywhere of a programme as a var statement
	const c = 100000

	//constant expressions gives arithmetic with arbitrary precision
	const x = 5e10 / c
	fmt.Println(x)

	//A numeric constant has no type until we give a type such as by an explicit conversion
	fmt.Println(int64(x))
	fmt.Println(float64(x))

	//A number can be given a type in a context such as a variable assignment or function call.
	fmt.Println(math.Cos(x))


}