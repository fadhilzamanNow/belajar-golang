package main

import (
	"fmt"
)

type Person struct {
	fname string
	lname string
	hobby []string
}




func main(){
	isBool := false
	angka := 12
	angkafloat := 12.1231
	fadhil := Person{
		fname : "Fadhil",
		lname : "Isfadhillah",
		hobby : []string{"Mancing","Ngoding","Makan"},
	}
	fmt.Printf("This is the value %+v dan angkanya adalah %T \n", fadhil, angka)
	fmt.Printf("The value of the %t is %v \n", isBool, isBool)
	fmt.Printf("The value of the float is %1.2f", angkafloat )

	fmt.Scanf("Are you good? ")
	fmt.Fscanf("Testing?")
}
