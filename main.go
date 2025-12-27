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

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
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

	canMakeIt(myEngine, myEngine.milesLeft())
	canMakeIt(electricEngine{20, 10}, 100)
}
