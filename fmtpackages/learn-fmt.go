package main

import (
	"fmt"
)

func main() {
	/* var name string
	var age int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write your name : ")

	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Terdapat error ketika memasukkan nama")
	}

	fmt.Printf("Hello %s \n", strings.TrimSpace(name))

	fmt.Print("What is your age : ")
	_, err2 := fmt.Scanln(&age)

	if err2 != nil {
		fmt.Println("Terdapat sebuah kesalahan lagi")
	}

	fmt.Printf("Your age is %d \n", age)

	confirmation := ""

	fmt.Print("Can you provide me your data back [yes/no] ? ")
	fmt.Scanf("%s", &confirmation)

	if strings.ToLower((confirmation)) == "yes" {
		fmt.Println("Okay thank you for your confirmation sir")
	} else if strings.ToLower((confirmation)) == "no" {
		fmt.Println("Sorry, if it disturb you")
	} */

	var answer string
	fmt.Sscan("Your", &answer)
	fmt.Println(answer)

	/* other1 := []byte{"Satu", "Dua", "Tiga"}

	fmt.Printf("Isi dari other adalah %v", other1) */

	/* 	s := make([]string, 5)
	   	d := []string{"Satu", "Dua", "Tiga"}
	   	sv := append(s, "Fadhil")
	   	d := append(d, "Isfadhillah")
	   	fmt.Printf("The value of s is %v and %v", sv, d) */

	s := make([]string, 0, 2)

	s = append(s, "ok")
	s = append(s, "ok2")
	s = append(s, "oke3", "nice")
	s = append(s, "oke2")
	s = append(s, "oke3", "oke4", "o", "o")
	fmt.Printf("The value of s is %v and the cap of s is %v but the length of s is %v", s, cap(s), len(s))
	data := []byte("Satu")
	sesuatu := "Tiga"
	data = fmt.Appendln(data, "This is also the new data")
	data = fmt.Appendf(data, "Sesuatu ya lah ges %v", sesuatu)

	fmt.Printf("Panjang dari datanya adalah %d dan isi dari datanya adalah %s", data, string(data))

}
