package main

import (
	"fmt"
)

type Person struct {
	fname string
	lname string
	hobby []string
}

func (p *Person) greeting() string {
	greet := fmt.Sprint("Hello guys, my name is ", p.fname, " ", p.lname, " and my hobby is ", p.hobby)
	return greet
}

func main() {
	isBool := false
	angka := 12
	angkafloat := 12.1231

	//initialize struct directly
	fadhil := Person{
		fname: "Fadhil",
		lname: "Isfadhillah",
		hobby: []string{"Mancing", "Ngoding", "Makan"},
	}
	fmt.Printf("This is the value %v dan angkanya adalah %T \n", &fadhil, angka)
	fmt.Printf("The value of the %t is %v \n", isBool, isBool)
	fmt.Printf("The value of the float is %1.2f", angkafloat)

	isfadhillah := &fadhil

	fmt.Printf("Memory pointer is %p \n", &fadhil)
	isfadhillah.fname = "Dhilz"

	fmt.Println("Fadhil informasinya adalah :", fadhil)

	// anonymous struct

	kucing := struct {
		name  string
		color string
		age   float64
	}{
		name:  "Joni",
		color: "Black",
		age:   5.4,
	}

	fmt.Println("Isi dari kucing adalah", kucing)

	mainan := struct {
		nama  string
		bahan string
	}{
		nama:  "Transformer",
		bahan: "Metal",
	}

	fmt.Println("Nama mainannya adalah", mainan)

	fmt.Println(fadhil.greeting())
}
