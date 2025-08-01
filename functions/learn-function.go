package main

import (
	"fmt"
	utilfunc "mainan-golang/functions/lib"
)

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Learn function")
	fmt.Println("===============")

	fmt.Println("The value of 3 + 7 is ", add(3, 7))

	fmt.Println("The value of 7 -3  is", utilfunc.Substract(7, 3))
	fmt.Println("The multiple of 2,3,4,5,6 is ", utilfunc.MultipleUltimate(2, 3, 4, 5, 6))

	_, secondVal := utilfunc.ReturnTwo(1, 2)
	fmt.Println("The return of the second values is : ", secondVal)

}
