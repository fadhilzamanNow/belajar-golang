package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	next := intSeq()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	aight := intSeq()

	fmt.Println(aight())
}
