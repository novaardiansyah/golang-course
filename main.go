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

	var intSlice2 []int32 = []int32{8, 9, 10}
	fmt.Printf("intSlice2: The length is %v with capacity %v\n", len(intSlice2), cap(intSlice2))

	intSlice2 = append(intSlice, intSlice2...)
	fmt.Printf("intSlice2: The length is %v with capacity %v\n", len(intSlice2), cap(intSlice2))
	fmt.Println(intSlice2)
}
