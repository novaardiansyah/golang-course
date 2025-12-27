package main

import (
	"fmt"
	"time"
)

func main() {
	var n uint32 = 10000000
	var testSlice = []int32{}
	var testSlice2 = make([]int32, 0, n)

	fmt.Printf("Total time without pre-allocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with pre-allocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int32, n uint32) time.Duration {
	var t0 = time.Now()
	for len(slice) < int(n) {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
