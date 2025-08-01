package main

import (
	"fmt"
)

func main() {
	randSlice := make([]int, 0)

	randSlice = append(randSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("=== RANGE OVER SLICE ====")
	for i, val := range randSlice {
		fmt.Println(i, val)
	}

	fmt.Println("==== RANGE OVER STRING ====")
	for i, val := range "Mantap" {
		fmt.Println(i, val)
	}

	randMap := map[string]string{
		"fname": "Fadhil",
		"lname": "Isfadhillah",
	}

	fmt.Println("==== RANGE OVER MAP ====")
	for i, val := range randMap {
		fmt.Println(i, val)
	}
}
