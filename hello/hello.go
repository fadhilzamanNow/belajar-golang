package main

import (
	"fmt"
	"mainan-golang/greeting"
	"log"
)

func main(){
	log.SetPrefix("greetings : ")
	log.SetFlags(0)

	message, err := greeting.Hello("");
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message);
}
