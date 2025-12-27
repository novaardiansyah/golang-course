package main

import (
	"fmt"
)

type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type owner struct {
	name string
	age uint8
}


func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func main() {
	var myEngine gasEngine = gasEngine{20, 10, owner{"John", 30}}
	fmt.Println(myEngine)

	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	}{20, 10}
	
	fmt.Println(myEngine2)

	fmt.Println(myEngine.milesLeft())
}
