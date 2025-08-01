package main

import (
	"fmt"
	"mainan-golang/array/lib"
)

func main() {
	fmt.Println("This is leanring array")
	intArrayNum := [3]int{1, 2, 3}
	hasil := lib.AddAllArray[int](intArrayNum)
	fmt.Println(hasil)

	//initialize empty array

	var emptyArray []string

	fmt.Printf("The empty array is %v \n", emptyArray)
	fmt.Println("The empty array is ", emptyArray, emptyArray == nil, len(emptyArray) == 0)

	//initialize empty slice

	emptySlice := make([]string, 5, 5)

	fmt.Println("The empty slice is ", emptySlice, "and the empty slice length is ", len(emptySlice), "and the capacity of the slice is ", cap(emptySlice))

	emptySlice = append(emptySlice, "First")

	fmt.Println("The empty slice is ", emptySlice, "and the length is ", len(emptySlice))

	copySlice := make([]string, len(emptySlice))

	fmt.Println("The copy slice is ", copySlice, "with length of ", len(copySlice))
	copySlice = append(copySlice, "Satu", "Dua", "Tiga", "Empat", "Lima")

	sliceSlice := copySlice[7:10]

	fmt.Println("Isi dari Slice 7-10 adalah ", sliceSlice, " with its length is ", len(sliceSlice))

	//initialize slice with populated data

	populatedSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("The value of the populated is ", populatedSlice, " with its length is ", len(populatedSlice))

	// multidimensional slices

	multiSlices := make([][]int, 3)

	fmt.Println("Isi dari multiSlices adalah :", multiSlices)

	for i := range 3 {
		innerLen := i + 1
		multiSlices[i] = make([]int, innerLen)

		for j := range innerLen {
			multiSlices[i][j] = i + j
		}
	}

	fmt.Println("Isi akhir dari multiSlices adalah : ", multiSlices)

	// multidimensional array

	var multiArray [5][4]string

	fmt.Println("Isi multidimenional Array : ", multiArray)
}
