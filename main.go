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

func main() {
	var myEngine gasEngine = gasEngine{20, 10, owner{"John", 30}}
	fmt.Println(myEngine)
}
