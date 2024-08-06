package main

import (
	"fmt"
	"math"
)

// a const or constants
const Statement string = "This is a constants"

func main() {
	fmt.Println(Statement)

	var statement string = "Hello world -> this is a variable"
	fmt.Println(statement)

	const num = 45

	fmt.Println("The Tan of number 45 is, ", math.Tan(num))

	var test string = "*"
	for i := 0; i < 4; i++ {
		test += "**"
		fmt.Print(test)
	}
}
