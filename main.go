package main

import (
	"fmt"
)

func main() {
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
}
