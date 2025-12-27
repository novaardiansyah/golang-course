package main

import (
	"fmt"
)

type gasEngine struct {
	mpg uint8
	gallons uint8
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 20, gallons: 10}
	fmt.Println(myEngine)
}
