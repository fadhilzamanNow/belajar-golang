package main

import (
	"fmt"
	"maps"
)

func main() {
	// initialize an empty map

	myOwnMap := make(map[string]int)

	myOwnMap["first"] = 1
	myOwnMap["second"] = 2

	fmt.Println("This is the map of my own ", myOwnMap)
	fmt.Printf("This is the map : %v", myOwnMap)

	// doesnt exist key

	fmt.Println("The value of map third : ", myOwnMap["now"], " and the length is ", len(myOwnMap))

	// delete key

	delete(myOwnMap, "second")

	fmt.Println("myOwn Map after being delete :", myOwnMap)

	clear(myOwnMap)
	fmt.Println("myOwnMap after being cleared : ", myOwnMap)

	myOwnMap["third"] = 3

	// return a key/value pair of a map

	// a returns if the value of the key, while b return a boolean that check whether the key exist or not
	a, b := myOwnMap["third"]

	fmt.Println("The value of a is ", a, " and the value of b is ", b)

	// initializing and declaring the value of a map directly

	directMap := map[string]int{"first": 1, "satu": 1}

	fmt.Println("The value of direct Map is : ", directMap)

	// check whether the two map is same or not

	if maps.Equal(myOwnMap, directMap) {
		fmt.Println("Oke sama")
	} else {
		fmt.Println("Beda kocak")
	}
}
