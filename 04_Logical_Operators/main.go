package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20
	var z int = 30

	fmt.Println(x < z && x > y)
	fmt.Println(x < z || x > y)
	fmt.Println(x != z || x <= y)

}