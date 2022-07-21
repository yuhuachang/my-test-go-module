package main

import (
	"fmt"

	"github.com/yuhuachang/my-test-go-module/mypack"
)

func main() {
	const a int = 3       // a is an int const
	var b int = 4         // b is an int variable
	b = 7                 // reassign value to b
	c := mypack.Sum(a, b) // c use implicit type
	fmt.Printf("%d + %d = %d\n", a, b, c)
}
