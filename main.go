package main

import (
	"fmt"
)

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 25, "Eve": 22, "Jane": 24, "Lucy": 23}
	fmt.Println(myMap2)

	fmt.Println(myMap2["John"])
	
	var age, ok = myMap2["John"]
	fmt.Println(age, ok)

	if ok {
		fmt.Println("John is", age)
	} else {
		fmt.Println("John is not in the map")
	}

	delete(myMap2, "Adam")
	fmt.Println(myMap2)

	for name := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, myMap2[name])
	}

	for i, v := range myMap2 {
		fmt.Printf("Key: %v, Value: %v\n", i, v)
	}
}
