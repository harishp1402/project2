package main

import (
	addition "add/logic"
	"fmt"
)

func main() {
	var a int
	var b int
	result := addition.Sum(a, b)
	fmt.Println("Sum of integers =", result)

}
