package main

import (
	"fmt"
	"mainan-golang/greeting"
	"log"
)

func main(){
	log.SetPrefix("greetings : ")
	log.SetFlags(0)

	listNama := []string{"Fadhil","Isfadhillah","Fatz",}
	message, err := greeting.Hellos(listNama);
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message);
}
