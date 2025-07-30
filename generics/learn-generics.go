package main

import (
	"fmt"
	// "mainan-golang/generics/lib"
	// "mainan-golang/generics/nogenerics"
	"mainan-golang/generics/withgenerics"
)

func main() {
	intMap := make(map[string]int64)
	intMap["first"] = 20	// "mainan-golang/generics/lib"

	intMap["second"] = 30

	floatMap := map[string]float64 {
		"first" : 20.5,
		"second" : 31.1,
	}

	fmt.Println("Int Result")
	// hasilInteger := nogenerics.SumInts(intMap)
	hasilInteger := withgenerics.SumIntOrFloat[string, int64](intMap)
	fmt.Println(hasilInteger)

	fmt.Println("Float Result")

	// hasilFloat := nogenerics.SumFloats(floatMap)
	hasilFloat := withgenerics.SumIntOrFloat[string, float64](floatMap)
	fmt.Println(hasilFloat)


}
