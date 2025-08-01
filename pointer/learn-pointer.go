package main

import (
	"fmt"
)

func editLiteral(a *int, b int) {
	*a = b
}

func main() {
	a := 20

	fmt.Println("The value of a before being edited is", a)

	editLiteral(&a, 30)

	fmt.Println("The value of a after being edited is", a)
}
